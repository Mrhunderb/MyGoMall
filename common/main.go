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
	r, err := c.Register(ctx, &pb.RegisterRequest{Username: "abc", Password: "123"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("id: %d, token: %s\n", r.Id, r.Token)
}
