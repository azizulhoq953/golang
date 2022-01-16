package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", ":8888") //ip :port

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}
	conn, err := nl.Accept() //layer -5 session ;layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr) //localhost ip[::1]:60941

	//byte
	conn.Write([]byte("Welcome To Our New Networking Code"))

	conn.Close()
	nl.Close()
}
