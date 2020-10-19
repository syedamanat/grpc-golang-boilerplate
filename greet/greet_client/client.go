package main

import (
	"fmt"
	"log"

	"github.com/syedamanat/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()
	//gets called at the very end because defer

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Created client %f", c)

}
