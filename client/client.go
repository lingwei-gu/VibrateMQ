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
	reply := &handler.String{Value: "Success" + args.GetValue()}
	fmt.Println("Success: " + args.GetValue())
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	handler.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	for {
		grpcServer.Serve(server)
	}

}
