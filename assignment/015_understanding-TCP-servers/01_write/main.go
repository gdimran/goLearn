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
			log.Fatalln(err)
			continue
		}

		io.WriteString(conn, "\nHello This is Imran From TCP server\n")
		fmt.Fprintln(conn, "How are you?")
		fmt.Fprintf(conn, "%v", "Hope you are doing well")

		conn.Close()
	}
}
