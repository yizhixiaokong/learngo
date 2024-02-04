package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Panicln("Listen tcp 127.0.0.1:8081 Err:", err)
	}

	defer server.Close()
	log.Println("HTTP Server Start, Listening 127.0.0.1:8081")
	reqID := 0
	for {
		client, err := server.Accept()
		if err != nil {
			log.Println("Accept Err:", err)
			continue
		}
		reqID += 1

		go func(client net.Conn, reqID int) {
			defer client.Close()
			lineReader := bufio.NewScanner(client)
			for lineReader.Scan() {
				text := lineReader.Text()
				if len(text) == 0 {
					break
				}
				log.Printf("requestID %d, Header Content:%s\n", reqID, text)
			}
			body := "<h1>Current Time: " + time.Now().Local().Format(time.RFC3339) + "</h1>"

			fmt.Fprint(client, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(client, "content-length:%d\r\n", len(body))
			fmt.Fprint(client, "Content-Type:text/html;charset=utf-8\r\n\r\n")
			fmt.Fprint(client, body)

		}(client, reqID)

	}
}
