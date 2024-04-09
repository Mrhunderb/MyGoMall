package user

import (
	"context"
	"fmt"
	"mygomall/common/db"
	"mygomall/service/user/endpoints"
	"mygomall/service/user/pb"
	"mygomall/service/user/service"
	"mygomall/service/user/transports"
	"mygomall/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var ServiceName = "user"

func RunUserService(port string) {

	var (
		userDB = viper.GetString("MySQL.User")
		passDB = viper.GetString("MySQL.Password")
		hostDB = viper.GetString("MySQL.Host")
		portDB = viper.GetString("MySQL.Port")
		nameDB = viper.GetString("MySQL.Database")
	)

	var (
		etcdHost = viper.GetString("Etcd.Host")
		etcdPort = viper.GetString("Etcd.Port")
	)

	logger := zap.NewExample()

	// Connect to MySQL
	db, err := db.Connect(userDB, passDB, hostDB, portDB, nameDB)
	if err != nil {
		panic(err)
	}
	logger.Info("database connected")

	svs := service.NewUserService(*logger, *db)
	endpoint := endpoints.MakeEndpoints(svs)
	grpcServer := transports.NewGRPCServer(endpoint)

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
		pb.RegisterUserServer(baseServer, grpcServer)
		logger.Info("server started", zap.String("address", fmt.Sprintf(":%s", port)))
		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Error("failed to server", zap.Error(err))
			errs <- err
		}
	}()

	logger.Info("exit", zap.Error(<-errs))
}
