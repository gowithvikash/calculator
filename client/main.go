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
	findSum(c)
	findPrimes(c)
	findAverage(c)
	findMaximum(c)
	findSquareRoot(c)
}

func findSum(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findSum Function Was Invoked At Client     _____")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 100, SecondNumber: 300})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("findSum Result: %v\n", res.Result)
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
		fmt.Printf("This is a prime : %v\n", res.Result)
	}

}
func findAverage(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findAverage Function Was Invoked At Client     _____")
}
func findMaximum(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findMaximum Function Was Invoked At Client     _____")
}
func findSquareRoot(c pb.CalculatorServiceClient) {
	fmt.Println("_____     findSquareRoot Function Was Invoked At Client     _____")
	res, err := c.SquareRoot(context.Background(), &pb.SqrtRequest{Number: 9})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res.Result: %v\n", res.Result)
}
