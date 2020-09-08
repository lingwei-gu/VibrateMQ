package connect

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// GetServerHost ...
func GetServerHost() (host string, err error) {
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty")
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

// GetConnect ...
func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// RegistServer ...
func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

// GetServerList ...
func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}

// SetServerNum ...
func SetServerNum(conn *zk.Conn, serverNum int) (stat *zk.Stat, err error) {
	_, stat, err = conn.Get("/go_servers")
	if err != nil {
		fmt.Println(err)
	}
	stat, err = conn.Set("/go_servers", []byte(strconv.Itoa(serverNum)), stat.Version)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// GetServerNum ...
func GetServerNum(conn *zk.Conn) (data []byte, stat *zk.Stat, err error) {
	data, stat, err = conn.Get("/go_servers")
	if err != nil {
		fmt.Println(err)
	}
	return
}
