package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello Imran This message from TCP Server\n")
		fmt.Fprintln(conn, "How are you?")
		fmt.Fprintln(conn, "%v", "Well I hope!")
		conn.Close()
	}
}
