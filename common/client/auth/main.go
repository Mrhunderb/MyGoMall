package main

import (
	"context"
	"fmt"
	"log"
	"mygomall/service/auth/pb"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	client, err := etcdv3.NewClient(context.Background(), []string{"localhost:2379"},
		etcdv3.ClientOptions{
			DialTimeout:   5 * time.Second,
			DialKeepAlive: 1 * time.Second,
		},
	)

	if err != nil {
		panic("etcdv3")
	}

	adds, err := client.GetEntries("/services/auth/")
	if err != nil {
		panic("etcdv3 get")
	}

	for _, addr := range adds {
		fmt.Println(addr)
	}

	conn, err := grpc.NewClient(adds[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetToken(ctx, &pb.GetTokenRequest{UserId: 1})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("token: %s", r.Token)
}
