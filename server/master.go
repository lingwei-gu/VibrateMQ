package main

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

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

func startServer(wg *sync.WaitGroup, port string) {
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

	wg.Done()

	for {
		grpcServer.Serve(server)
	}
}

func updateServerRecords() {
	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()

	// retrieve server list
	serverList, err := cnt.GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	// count how many servers it has
	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty")
		return
	}

	// update the data in znode
	_, err = cnt.SetServerNum(conn, count)
	if err != nil {
		fmt.Printf(" set znode server number error: %s \n", err)
		return
	}
}

func retrieveServerRecords() {
	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	data, _, err := cnt.GetServerNum(conn)

	fmt.Printf("Number of servers: %s \n", string(data))
}

func main() {
	var wg sync.WaitGroup
	for i := 8081; i <= 8083; i++ {
		wg.Add(1)
		go startServer(&wg, strconv.Itoa(i))
	}

	wg.Wait()

	updateServerRecords()
	retrieveServerRecords()
	hold := make(chan bool, 1)
	<-hold

	fmt.Printf("Terminated")
}
