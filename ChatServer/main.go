package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", ":8888")
	if err != nil {

		fmt.Println(err.Error())
		os.Exit(1) //(1) stop With error
	}
	defer nl.Close() //awaite
	log.Printf("Server Started on:8888")
	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			//continue

		}

		bs := make([]byte, 1024) //text
		n, e := conn.Read(bs)
		if e != nil {
			fmt.Println(e.Error())
		}
		fmt.Println(n)

		reqestr := string(bs) //convertion
		fmt.Println(reqestr)
		msg := fmt.Sprintf(`Your Message:%s`, reqestr)
		conn.Write([]byte(msg))
	}

}
