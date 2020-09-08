package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	cnt "VibrateMQ/connection"
	"VibrateMQ/handler"

	"google.golang.org/grpc"
)

// getServerHost ...
func getServerHost() (host string, err error) {
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := cnt.GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

func main() {
	serverHost, err := getServerHost()
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
