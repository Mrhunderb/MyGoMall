package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
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

	adds, err := client.GetEntries("/services/user/")
	if err != nil {
		panic("etcdv3 get")
	}

	for _, addr := range adds {
		fmt.Println(addr)
	}

	// conn, err := grpc.NewClient(adds[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	// defer conn.Close()
	// c := pb.NewUserClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.UserInfo(ctx, &pb.UserInfoRequest{Id: 1})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// log.Printf("id: %d, name: %s\n", r.Id, r.Username)
}
