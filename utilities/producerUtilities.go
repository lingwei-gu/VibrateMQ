package utilities

import (
	cnt "VibrateMQ/connection"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// GetServerPort gets the port number as a client-side service
func GetServerPort(index int) (serverHost string) {
	serverHost, err := GetServerHost(index)
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	fmt.Println("connect host: " + serverHost)
	return
}

// GetConns returns a list of connections with the worker server
func GetConns() (conns []*grpc.ClientConn, ports []string) {
	ports = GetAllHosts()

	len := GetServerLen(cnt.ServerPath)
	fmt.Printf("Server Number: %d \n", len)
	for i := 0; i < len; i++ {
		conn, err := grpc.Dial("localhost:"+GetServerPort(i), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		conns = append(conns, conn)
	}
	return
}
