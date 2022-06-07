package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("Server Is Failing To Listen")
	}
	fmt.Println("Server Is Listning On Adderss : ", address)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal("Server Is Failing To Serve Client")
	}

}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Println("______     Sum Function Was Invoked At Server     _____")
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	fmt.Println("______     Primes Function Was Invoked At Server     _____")
	var num int64 = in.Number
	var div int64 = 2

	for num > 1 {
		if num%div == 0 {
			err := stream.Send(&pb.PrimesResponse{Result: div})
			if err != nil {
				log.Fatal(err)
			}
			num /= div
		} else {
			div++
		}

	}

	return nil
}
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	fmt.Println("______     Average Function Was Invoked At Server     _____")
	var total = 0
	var count = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(total) / float64(count),
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		total += int(req.Number)
		count++
	}
}

func (s *Server) Maximum(stream pb.CalculatorService_MaximumServer) error {
	fmt.Println("______     Maximum Function Was Invoked At Server     _____")
	var max = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		if req.Number > int64(max) {
			err = stream.Send(&pb.MaximumResponse{
				Result: req.Number,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		max = int(req.Number)

	}
}
func (s *Server) SquareRoot(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	fmt.Println("______     SquareRoot Function Was Invoked At Server     _____")
	return &pb.SqrtResponse{Result: math.Sqrt(float64(in.Number))}, nil
}
