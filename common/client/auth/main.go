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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDE3MTI2NzI3OTksImlhdCI6MTcxMjY3Mjc5OSwidWlkIjoxfQ.-zzv5t8NbfMPkpKi1z5ahnPcl7aGPUmlffWs_pa3H0U"
	// r, err := c.GetToken(ctx, &pb.GetTokenRequest{UserId: 1})
	r, err := c.VerifyToken(ctx, &pb.Token{Token: token})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("valid: %v", r.Valid)
}
