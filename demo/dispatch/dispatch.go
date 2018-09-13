package main

import (
	"io"
	"log"
	"net"
)

func main() {

	addr := "0.0.0.0:8001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Failed to net.Listen(%s,%s),err=%v\n", "tcp", addr, err)
		return
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Failed to accept,err=%v", err)
			continue
		}
		go dispatch(conn)
	}
}

func dispatch(conn net.Conn) {
	buf := make([]byte, 1024*2)
	for {
		_, err := conn.Read(buf)
		if err == io.EOF {
			break
		}

		log.Printf("Successed to read,data=%v", string(buf))
	}
}
