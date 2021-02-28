package main

import (
	"fmt"
	"net"
	"os"

)

func main() {
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	conn.Write([]byte("Welcome we have received your message Imran"))

	conn.Close()
	nl.Close()
}
