package main

import (
	"com.grpc.zhand/calculatorEndterm/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculateServiceClient(conn)
	doManyTimesFromServer(c)
}

func doManyTimesFromServer(c calculatorpb.CalculateServiceClient) {
	ctx := context.Background()
	req := &calculatorpb.CalculatorRequest{Calculating: &calculatorpb.Calculating{
		Number: 120,
	}}
	stream, err := c.CalculateManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("error while calling CalculateManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error while reciving from CalculateManyTimes RPC %v", err)
		}
		log.Printf("response from CalculateManyTimes:%v \n", res.GetResult())
	}

}

func computeAverage(c calculatorpb.CalculateServiceClient) {

	requests := []*calculatorpb.AverageRequest{
		{
			Numbers: 1,
		},
		{
			Numbers: 2,
		},
		{
			Numbers: 3,
		},
		{
			Numbers: 4,
		},
	}

	ctx := context.Background()
	stream, err := c.ComputeAverage(ctx)
	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("Compute Average Response: %v\n", res)
}
