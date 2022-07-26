package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", "Localhost:8888") //port 1 to 65535
	if err != nil {

		fmt.Println(err.Error())		
		os.Exit(1) // (1) stop with Error

	}

	defer nl.Close() //awaite
	log.Printf("server Started on :8888")

	for {
		conn, err := nl.Accept()
		if err != nil {

			fmt.Println(err.Error())
			//continue

		}

		//scanner := bufio.NewScanner(com)
		//fmt.Println(line)
		//if line ==""{
		//	fmt.Println("line Ends")
		//	break
		// }

		bs := make([]byte, 1024) //text assa byte  to string
		n, e := conn.Read(bs)
		if e != nil {

			fmt.Println(e.Error())

		}
		fmt.Println(n)
		//fmt.Println(bs)
		reqstr := string(bs) //convertion
		fmt.Println(reqstr)

		//reqslc := string.Field(reqstr)
		//fmt.Println(reqslc,len(reqslc))

		body := `<html>
		<body>
		<title> My First Server In Go  </title>
		<h1> Server In GO Bangladesh </h1>

		<h2> Practice For Coading BootCamp For Mostain Billah Sir</h2>


		</body>
		</html>`

		resp := "HTTP/1.1 200 ok\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n%s"

		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)
		conn.Write([]byte(msg))
	}

}
