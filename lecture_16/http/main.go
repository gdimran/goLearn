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
	}

	bs := make([]byte, 1024)

	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n)
	reqster := string(bs) //converting bite to string format
	fmt.Println(reqster)

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Page Title</title></head><body><h1>Hello Imran</h1><p>This is a working.</p></body></html>`

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"
	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))

}
