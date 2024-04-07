package main

import (
	"context"
	"log"
	"mygomall/service/user/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:55055"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserInfo(ctx, &pb.UserInfoRequest{Id: 1})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("id: %d, name: %s\n", r.Id, r.Username)
}
