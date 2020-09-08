package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	cnt "VibrateMQ/connection"
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

func startServer(port string) {
	grpcServer := grpc.NewServer()
	handler.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()

	// register znode
	err = cnt.RegistServer(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		grpcServer.Serve(server)
	}
}

func main() {
	go startServer("8083")
	go startServer("8081")
	startServer("8082")
	fmt.Printf("success")
}
