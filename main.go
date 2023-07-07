package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	host := "0.0.0.0"
	port := "23"
	Type := "tcp"
	listen, err := net.Listen(Type, host+":"+port)
	if err != nil {
		log.Fatal("listen err:", err)
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("listen.Accept failed:", err)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//buffer := make([]byte, 1024)
	//
	//_, err := conn.Read(buffer)
	//if err != nil {
	//	return
	//}
	remote := conn.RemoteAddr().String()
	lastColon := strings.LastIndex(remote, ":")
	//parts := strings.
	ip := remote[0:lastColon]
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05")+" visit from:", ip)
	conn.Write([]byte("you ip is " + ip + "\n"))
	conn.Close()
}
