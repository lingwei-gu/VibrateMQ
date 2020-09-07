package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"VibrateMQ/handler"

	grpc "google.golang.org/grpc"
)

// HelloServiceImpl ...
type HelloServiceImpl struct{}

// Hello ...
func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *handler.String,
) (*handler.String, error) {
	reply := &handler.String{Value: "hello:" + args.GetValue()}
	go dial(args.GetValue())
	return reply, nil
}

func dial(args string) {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := handler.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &handler.String{Value: args})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}

func main() {
	grpcServer := grpc.NewServer()
	handler.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	for {
		grpcServer.Serve(server)
	}

}
