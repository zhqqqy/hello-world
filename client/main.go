package main

import (
	"google.golang.org/grpc"
	"log"
	pb "hello-world/proto"
	"os"
	"context"
	"time"
)

const (
	address  = "localhost:50051"
	defaultName = "worldaaa"
)

type Config struct {
	IP string `json:"ip"`
	Port string `json:"port"`
}

func main() {

conn,err:=grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name:=defaultName
	if len(os.Args)>1{
		name=os.Args[1]
	}

	for{
		r,err:=c.SayHello(context.Background(),&pb.HelloRequest{Name:name})
		if err!=nil{
			//code := grpc.Code(err)
			log.Println("could not greet: %v", err)
			continue
		}
		log.Printf("Greeting: %s", r.Message)
		time.Sleep(time.Second)
	}
}