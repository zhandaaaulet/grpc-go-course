package main

import (
	"com.grpc.zhand/calculator/calculatorpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type Server struct {
	calculatorpb.UnimplementedCalculateServiceServer
}

func (s *Server) CalculateEveryone(stream calculatorpb.CalculateService_CalculateEveryoneServer) error {
	fmt.Printf("CalculateEveryone function was invoked with a streaming requset\n")
	var max = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err.Error())
			return err
		}

		number := req.GetCalculating().GetNumber()
		if number > int64(max) {
			max = int(number)
		}
		if number < int64(max) {
			continue
		}
		result := "Numbers: " + strconv.FormatInt(int64(max), 10)
		err = stream.Send(&calculatorpb.CalculatorResponse{Result: result})
		if err != nil {
			log.Fatalf("error while sending to client: %v", err.Error())
			return err
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
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
