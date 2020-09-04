package main

import (
	"log"
	"net"
	"net/rpc"
)

// HelloService ...
type HelloService struct{}

// Hello ...
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
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
