package utilities

import "fmt"

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

func MaintainConns() {

}
