package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cnt "VibrateMQ/connection"
	"VibrateMQ/handler"

	"google.golang.org/grpc"
)

func main() {
	serverHost, err := cnt.GetServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	fmt.Println("connect host: " + serverHost)
	conn, err := grpc.Dial("localhost:"+serverHost, grpc.WithInsecure())
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
