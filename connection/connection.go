package connect

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

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
