package main

import (
	"bufio"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
)

var flag string

var pubKeyX *big.Int
var pubKeyY *big.Int

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

	pubKeyX, _ = new(big.Int).SetString("f2a16842f109af6c5c45677cec992e9893161872646ed2fd270a9069234f2d45", 16)
	pubKeyY, _ = new(big.Int).SetString("74c9bd8482f05599502601328cecd7253f4941eac633f972b72f82ca7466ca5e", 16)
}

func readHex(scanner *bufio.Scanner) *big.Int {
	rString := scanner.Text()
	r, ok := new(big.Int).SetString(rString, 16)
	if !ok {
		return nil
	}
	return r
}

// hashToInt converts a hash value to an integer. Per FIPS 186-4, Section 6.4,
// we use the left-most bits of the hash to match the bit-length of the order of
// the curve. This also performs Step 5 of SEC 1, Version 2.0, Section 4.1.3.
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

// https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm
func checkSignature(r, s *big.Int, challenge []byte) bool {
	curve := elliptic.P256()

	e := hashToInt(challenge[:], curve)
	N := curve.Params().N

	// sInv := new(big.Int).ModInverse(s, N)
	// Optimize modular inverse calculation by using Fermat's little theorem
	// Modular inverse of s, sInv is number that s * sInv = 1 (mod N)
	// Fermat's little theorem:
	// x^(n - 1) = x^0 = 1
	// x^(n - 2) = x^(-1) because x^(n - 2) * x = x^(n - 1) which is 1
	sInv := new(big.Int).Exp(s, new(big.Int).Sub(N, big.NewInt(2)), N)

	// u1 = e * s^(-1) mod N
	u1 := e.Mul(e, sInv)
	u1.Mod(u1, N)

	// u2 = r * s^(-1) mod N
	u2 := sInv.Mul(r, sInv)
	u2.Mod(u2, N)

	u1GX, u1GY := curve.ScalarBaseMult(u1.Bytes())
	u2QaX, u2QaY := curve.ScalarMult(pubKeyX, pubKeyY, u2.Bytes())
	// u1 * G + u2 * PubKey
	x, _ := curve.Add(u1GX, u1GY, u2QaX, u2QaY)
	x.Mod(x, N)
	return x.Cmp(r) == 0
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Hi!\n"))

	challenge := make([]byte, 32)
	n, err := rand.Read(challenge)
	if err != nil {
		log.Printf("Failed to generate challenge: %v", err)
		return
	}
	if n != len(challenge) {
		log.Printf("Failed to generate challenge: not enough entropy gathered")
		return
	}

	scanner := bufio.NewScanner(conn)

	conn.Write([]byte(fmt.Sprintf("Please sign this data: %x\n", challenge)))

	conn.Write([]byte("r: "))
	if ok := scanner.Scan(); !ok {
		return
	}
	r := readHex(scanner)
	if r == nil {
		conn.Write([]byte("Invalid r! (╯°□°）╯︵ ┻━┻\n"))
		return
	}

	conn.Write([]byte("s: "))
	if ok := scanner.Scan(); !ok {
		return
	}
	s := readHex(scanner)
	if s == nil {
		conn.Write([]byte("Invalid s! ح(•̀ж•́)ง\n"))
		return
	}

	if checkSignature(r, s, challenge) {
		conn.Write([]byte(fmt.Sprintf("Signature verified! Your flag is: %s\n", flag)))
	} else {
		conn.Write([]byte("Bad signature! {ಠʖಠ}"))
	}
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
