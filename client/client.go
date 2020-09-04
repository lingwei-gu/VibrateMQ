package main

import (
	"log"
	"net"
	"net/rpc"
)

// ReceiveService ...
type ReceiveService struct{}

// Receive ...
func (p *ReceiveService) Receive(request string, reply *string) error {
	*reply = "Receive:" + request
	return nil
}

func main() {
	rpc.RegisterName("ReceiveService", new(ReceiveService))
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
			continue
		}
		rpc.ServeConn(conn)
	}

}
