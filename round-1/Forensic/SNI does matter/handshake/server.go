package main

import (
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Hi :))\nFlag is: HZU18{TLS_0R_SSL?_21ec448ad6d3}"))
}

func main() {
	var port string = "8080"
	if portStr, ok := os.LookupEnv("PORT"); ok {
		port = portStr
	}
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
