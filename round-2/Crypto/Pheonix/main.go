package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
)

var (
	privKeyD *big.Int
	pubKeyX  *big.Int
	pubKeyY  *big.Int
)

func init() {
	privKeyD, _ = new(big.Int).SetString("<REDACTED>", 16)
	pubKeyX, _ = new(big.Int).SetString("f2a16842f109af6c5c45677cec992e9893161872646ed2fd270a9069234f2d45", 16)
	pubKeyY, _ = new(big.Int).SetString("74c9bd8482f05599502601328cecd7253f4941eac633f972b72f82ca7466ca5e", 16)
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

var (
	curve = elliptic.P256()
	N     = curve.Params().N
)

func secureRandom() *big.Int {
	// TODO
	return big.NewInt(4)
}

func fermatInverse(k, N *big.Int) *big.Int {
	two := big.NewInt(2)
	nMinus2 := new(big.Int).Sub(N, two)
	return new(big.Int).Exp(k, nMinus2, N)
}
func sign(privD *big.Int, data []byte) (r *big.Int, s *big.Int) {
	hash := sha256.Sum256(data)
	e := hashToInt(hash[:], curve)

	k := secureRandom()
	kInv := fermatInverse(k, N)

	r, _ = curve.ScalarBaseMult(k.Bytes())
	r.Mod(r, N)

	s = new(big.Int).Mul(privD, r)
	s.Add(s, e)
	s.Mul(s, kInv)
	s.Mod(s, N)
	return
}

func main() {
	r1, s1 := sign(privKeyD, []byte("hello world"))
	println(r1.String(), s1.String())
	r2, s2 := sign(privKeyD, []byte("hi world"))
	println(r2.String(), s2.String())
}
