package main

import (
	"com.grpc.zhand/calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculateServiceClient(conn)
	doCalculatingEveryone(c)
}

func doCalculatingEveryone(c calculatorpb.CalculateServiceClient) {

	stream, err := c.CalculateEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while open stream: %v", err)
	}

	requests := []*calculatorpb.CalculatorRequest{
		{
			Calculating: &calculatorpb.Calculating{
				Number: 1,
			},
		},
		{
			Calculating: &calculatorpb.Calculating{
				Number: 5,
			},
		},
		{
			Calculating: &calculatorpb.Calculating{
				Number: 3,
			},
		},
		{
			Calculating: &calculatorpb.Calculating{
				Number: 6,
			},
		},
		{
			Calculating: &calculatorpb.Calculating{
				Number: 2,
			},
		},
		{
			Calculating: &calculatorpb.Calculating{
				Number: 20,
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending req to server: %v", err.Error())
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("errow while closing client's stream")
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while getting message from stream: %v", err.Error())
			}
			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()


	<-waitc
}