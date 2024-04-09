package auth

import (
	"context"
	"fmt"
	"mygomall/service/auth/endpoints"
	"mygomall/service/auth/pb"
	"mygomall/service/auth/service"
	"mygomall/service/auth/transports"
	"mygomall/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var ServiceName = "auth"

func RunAuthService(port string) {

	var (
		redisHost = viper.GetString("Redis.Host")
		redisPort = viper.GetString("Redis.Port")
		redisPass = viper.GetString("Redis.Password")
	)

	var (
		etcdHost = viper.GetString("Etcd.Host")
		etcdPort = viper.GetString("Etcd.Port")
	)

	logger := zap.NewExample()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	})
	logger.Info("redis connected")

	// Create service
	svs := service.NewAuthService(*logger, *rdb)
	endpoint := endpoints.MakeEndpoints(svs)
	grpcServer := transports.NewAuthGRPCServer(endpoint)

	// Connect to etcd
	etdcClient, err := etcdv3.NewClient(
		context.Background(),
		[]string{etcdHost + ":" + etcdPort},
		etcdv3.ClientOptions{
			DialTimeout:   time.Second * 3,
			DialKeepAlive: time.Second * 3,
		},
	)
	if err != nil {
		panic(err)
	}
	ip, _ := utils.GetOutboundIP()

	// Register service to etcd
	register := etcdv3.NewRegistrar(etdcClient, etcdv3.Service{
		Key:   fmt.Sprintf("/services/%s/%s:%s", ServiceName, ip, port),
		Value: fmt.Sprintf("%s:%s", ip, port),
	}, log.NewJSONLogger(os.Stdout))

	register.Register()
	defer register.Deregister()
	logger.Info("service registered")

	// Graceful shutdown
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Start gRPC server
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Error("failed to listen", zap.String("address", fmt.Sprintf(":%s", port)))
		os.Exit(1)
	}
	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterAuthServiceServer(baseServer, grpcServer)
		logger.Info("server started", zap.String("address", fmt.Sprintf(":%s", port)))
		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Error("failed to server", zap.Error(err))
			errs <- err
		}
	}()

	logger.Info("exit", zap.Error(<-errs))
}
