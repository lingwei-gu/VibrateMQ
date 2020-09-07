package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"VibrateMQ/handler"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := handler.NewHelloServiceClient(conn)

	for i := 0; i < 1000; i++ {
		reply, err := client.Hello(context.Background(), &handler.String{Value: strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue(), i)
	}
}
