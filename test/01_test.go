package test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"
)

func Test01(test *testing.T) {
	t := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(t)
	//fmt.Println(time.Now().Unix())
	//s:="你好"
	b := []byte(t)
	fmt.Println(b)
	h := bytes.Join([][]byte{[]byte("123"), []byte("123")}, []byte{})
	fmt.Println(h)
}

func TestSha256(test *testing.T) {
	s := "123456432423"
	ss := sha256.Sum256([]byte(s))
	fmt.Printf("%x", ss)
}

func TestLsh(test *testing.T) {
	target := big.NewInt(1)
	fmt.Println(target)
	target.Lsh(target, uint(232))
	fmt.Println(target)
}

func TestBig(test *testing.T) {
	var h big.Int
	h.SetBytes([]byte("34321321321dxsadd"))
	fmt.Println(h)
}

func TestWY(test *testing.T) {
	var hashInt big.Int
	var hash [32]byte
	target := big.NewInt(1)
	target.Lsh(target, uint(256-10))
	fmt.Println(target)
	s := "I love you"
	var n int64 = 1
	for {
		s += string(n)
		hash = sha256.Sum256([]byte(s))
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			fmt.Println("success:", hashInt.String())
			fmt.Printf("%x", hash)
			fmt.Println()
			fmt.Println("n:", n)
			break
		} else {
			n++
		}
	}
	fmt.Printf("%x", hash)
}
