package main

import (
	"com.grpc.zhand/calculatorEndterm/calculatorpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type Server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

//PrimeNumberDecomposition
func (s *Server) CalculateManyTimes(req *calculatorpb.CalculatorRequest, stream calculatorpb.CalculateService_CalculateManyTimesServer) error {
	fmt.Printf("CalculateManyTimes function was invoked with %v \n", req)
	number := req.GetCalculating().GetNumber()
	n := int(number)
	arr := PrimeNumberDecomposition(n)
	for i := 0; i < len(arr); i++ {
		res := &calculatorpb.CalculatorResponse{Result: fmt.Sprintf("%d || Server response: %v\n", n, arr[i])}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending calculate many times responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

//Logic of prime number decomposition
func PrimeNumberDecomposition(n int) (arr []int) {
	for n%2 == 0 {
		arr = append(arr, 2)
		n /= 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			arr = append(arr, i)
			n /= i
		}
	}
	if n > 2 {
		arr = append(arr, n)
	}
	return
}

//ComputeAverage
func (s *Server) ComputeAverage(stream calculatorpb.CalculateService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a streaming request\n")
	var result float64
	var arr []int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result = FindAverage(arr)
			return stream.SendAndClose(&calculatorpb.AverageResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		arr = append(arr, req.GetNumbers())
	}
}

//The logic of finding the average
func FindAverage(arr []int64) float64 {
	sum := 0.0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	avg := sum / float64(n)
	return avg
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculateServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
