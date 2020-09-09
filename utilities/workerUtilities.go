package utilities

import (
	cnt "VibrateMQ/connection"
	"errors"
	"fmt"
	"strconv"
)

// AddZnode adds a new znode by the given path
func AddZnode(path string) {
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()

	err = cnt.RegistNode(conn, path)
	if err != nil {
		fmt.Printf("Register new znode error: %s \n", err)
	}
}

// GetServerHost gets the server according to the load balancer
func GetServerHost(index int) (host string, err error) {
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()

	serverList, err := cnt.GetServerList(conn, cnt.ServerPath)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	host = serverList[index]

	// balancerIndex := RetrieveRecords(cnt.BalancerPath)
	// host = serverList[balancerIndex]
	return
}

// GetServerLen counts how many children are in the current path
func GetServerLen(path string) (count int) {
	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()

	// retrieve server list
	serverList, err := cnt.GetServerList(conn, path)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	// count how many servers it has
	count = len(serverList)
	if count == 0 {
		err = errors.New("server list is empty")
		return
	}
	return
}

// UpdateRecords updates the int data in the znode of given path
func UpdateRecords(path string, count int) {
	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	// update the data in znode
	_, err = cnt.SetServerNum(conn, count, path)
	if err != nil {
		fmt.Printf(" set znode server number error: %s \n", err)
		return
	}
}

// RetrieveRecords retrieves the int data from the znode of given path
func RetrieveRecords(path string) (serverNum int) {
	// connect to zk
	conn, err := cnt.GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	data, _, err := cnt.GetServerNum(conn, path)

	fmt.Printf("Number of servers: %s \n", string(data))
	serverNum, err = strconv.Atoi(string(data))
	if err != nil {
		fmt.Printf("string to int error: %s \n", err)
	}
	return serverNum
}
