package connect

import (
	"fmt"
	"strconv"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// GetConnect connects to zk server
func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// RegistServer registers servers
func RegistServer(conn *zk.Conn, host string, path string) (err error) {
	_, err = conn.Create(path+"/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

// RegistNode registers znode by path
func RegistNode(conn *zk.Conn, path string) (err error) {
	_, err = conn.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

// GetServerList gets all the children servers under the path
func GetServerList(conn *zk.Conn, path string) (list []string, err error) {
	list, _, err = conn.Children(path)
	return
}

// SetServerNum sets the number of its children servers
func SetServerNum(conn *zk.Conn, serverNum int, path string) (stat *zk.Stat, err error) {
	_, stat, err = conn.Get(path)
	if err != nil {
		fmt.Println(err)
	}
	stat, err = conn.Set(path, []byte(strconv.Itoa(serverNum)), stat.Version)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// GetServerNum gets the number of its children servers
func GetServerNum(conn *zk.Conn, path string) (data []byte, stat *zk.Stat, err error) {
	data, stat, err = conn.Get(path)
	if err != nil {
		fmt.Println(err)
	}
	return
}
