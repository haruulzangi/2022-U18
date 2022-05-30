package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	a = 4332
	b = 322
	n = 1812433253
)

func rand(state *int) {
	*state = (*state*a + b) % n
}

var flag string

func init() {
	flagFile, err := os.Open("flag.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(flagFile)
	if !scanner.Scan() {
		panic("Failed to read flag from file")
	}
	flag = scanner.Text()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	state := 3213921
	for i := 0; i < 5; i++ {
		rand(&state)
		conn.Write([]byte("Би ямар тоо таасан бэ? "))
		if ok := scanner.Scan(); !ok {
			return
		}

		guess, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if guess != state {
			conn.Write([]byte("Худлаа яриад байгаарай!\n"))
			return
		}
	}

	conn.Write([]byte(fmt.Sprintf("Nice! Туг: %s\n", flag)))
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
			panic(err)
		}
		go handleConnection(conn)
	}
}
