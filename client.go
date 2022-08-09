package main

import (
	"context"
	"flag"
	"log"

	"github.com/itsapep/calculator-grpc-client/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create access using flag
	number1 := flag.Int("num1", 0, "Number 1")
	number2 := flag.Int("num2", 1, "Number 1")
	operator := flag.String("opr", "+", "Calculation Operator (+,-,*,/)")
	serverHost := flag.String("srv", "localhost:8888", "Server Host")

	flag.Parse()
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(*serverHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect ...", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	c := api.NewCalculatorClient(conn)

	response, err := c.DoCalc(context.Background(), &api.CalculatorInputMessage{
		Number1:  int32(*number1),
		Number2:  int32(*number2),
		Operator: *operator,
	})

	if err != nil {
		log.Fatalln("error when calling ...", err)
	}

	log.Printf("the result is %v", response.ResNumber)
}
