package test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
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
