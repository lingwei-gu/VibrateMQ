package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"VibrateMQ/handler"
	util "VibrateMQ/utilities"
)

func main() {
	connections, ports := util.GetConns()
	connNum := len(connections)
	// for to make clients
	var clients []handler.HelloServiceClient
	for _, conn := range connections {
		clients = append(clients, handler.NewHelloServiceClient(conn))
	}

	fmt.Printf("connNum: %d\n", connNum)
	// to check numGoroutines, use a seperate goroutine to check with time.sleep until the following for loop ends
	// if too many goroutines, add servers and update connections, and vice versa

	for i := 0; i < 1000; i++ {
		reply, err := clients[i%connNum].Hello(context.Background(), &handler.String{Value: strconv.Itoa(i)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue()+" ", i, "\nPort Used: ", ports[i%connNum])
	}
	for _, conn := range connections {
		defer conn.Close()
	}
}
