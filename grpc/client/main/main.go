package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/server"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn fail, %v", err)
	}
	defer conn.Close()
	client, _ := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "fuzicheng"})
	fmt.Printf("%v", resp.GetResponseMsg())
}
