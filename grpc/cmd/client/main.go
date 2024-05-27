package main

import (
	"context"
	"grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloClient(conn)

	req := pb.HelloRequest{Name: "Gopher"}

	res, err := client.SayHello(context.Background(), &req)

	if err != nil {
		panic(err)
	}

	log.Println(res)

}
