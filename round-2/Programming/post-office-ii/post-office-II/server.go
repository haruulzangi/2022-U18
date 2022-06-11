package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
)

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

	tries := uint(512)
	max := new(big.Int).Lsh(big.NewInt(2), tries-1)

	cell, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Printf("Failed to generate random cell: %v", err)
		return
	}
	cell = new(big.Int).Add(cell, big.NewInt(1))

	scanner := bufio.NewScanner(conn)
	conn.Write([]byte("Шуудангийн салбарт тавтай морил!\n"))
	conn.Write([]byte(fmt.Sprintf("Хайрцагуудаа 1-%s хүртэл дугаарласан байгаа, тугтай хайрцагийг олоорой!\n", max.String())))
	conn.Write([]byte(fmt.Sprintf("%d хайрцагийг нээх боломжтой байгаа шүү ʘ‿ʘ\n", tries)))
	for i := uint(0); i < tries; i++ {
		conn.Write([]byte("Хайрцагны дугаар: "))
		if !scanner.Scan() {
			return
		}
		guess, ok := new(big.Int).SetString(scanner.Text(), 10)
		if !ok {
			conn.Write([]byte("Энэ тоо байх ёстой юу? ¿ⓧ_ⓧﮌ\n"))
			return
		}

		if guess.Sign() <= 0 || guess.Cmp(max) == 1 {
			conn.Write([]byte("Энэ юу вэ? (⊙.☉)7\n"))
			return
		}

		order := cell.Cmp(guess)
		if order == 0 {
			conn.Write([]byte(fmt.Sprintf("Nice! Туг: %s\n", flag)))
			return
		}

		if order < 0 {
			conn.Write([]byte("Хэтрээд явчихав уу?\n"))
		} else {
			conn.Write([]byte("Арай наана уу?\n"))
		}
	}

	conn.Write([]byte("¯\\_(⊙︿⊙)_/¯"))
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
