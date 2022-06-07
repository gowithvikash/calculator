package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	fmt.Println("_____________________________________________________________________")
	findSum(c)
	fmt.Println("_____________________________________________________________________")
	findSquareRoot(c)
	fmt.Println("_____________________________________________________________________")
	findPrimes(c)
	fmt.Println("_____________________________________________________________________")
	findAverage(c)
	fmt.Println("_____________________________________________________________________")
	findMaximum(c)
	fmt.Println("_____________________________________________________________________")

}

func findSum(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findSum Function Was Invoked At Client     _____")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 100, SecondNumber: 300})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("_____findSum Result : %v\n", res.Result)

}
func findSquareRoot(c pb.CalculatorServiceClient) {
	fmt.Println("_____    findSquareRoot Function Was Invoked At Client     _____")
	res, err := c.SquareRoot(context.Background(), &pb.SqrtRequest{Number: 9})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("_____    res.Result : %v\n", res.Result)
	fmt.Println("")
}
func findPrimes(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findPrimes Function Was Invoked At Client     _____")

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{Number: 48})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("____  this is a prime : %v\n", res.Result)
	}

}
func findAverage(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findAverage Function Was Invoked At Client     _____")
	var reqs = []*pb.AverageRequest{{Number: 100}, {Number: 55}, {Number: 55}, {Number: 67}, {Number: 78}, {Number: 71}, {Number: 45}}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range reqs {
		stream.Send(v)
	}
	stream.CloseSend()
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("_____    findAverage_Result : %v\n", res.Result)

}

func findMaximum(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findMaximum Function Was Invoked At Client     _____")
	fmt.Println("")
	var requests = []*pb.MaximumRequest{{Number: 7}, {Number: 8}, {Number: 1}, {Number: 25}, {Number: 14}, {Number: 6}, {Number: 88}}
	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	waitc := make(chan struct{})

	go func() {
		for _, v := range requests {
			if err = stream.Send(v); err != nil {
				log.Fatal(err)
			}

		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("This Is Max : %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
