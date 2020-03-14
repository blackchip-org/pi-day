package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	// piServer = "127.0.0.1"
	piServer = "192.168.7.77"
	inPort   = 3141
	outPort  = 31415
)

var lut []string

func reader(out *bufio.Writer) {
	host := fmt.Sprintf("%v:%v", piServer, inPort)
	fmt.Printf("connecting to %v\n", host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalf("error connecting to server: %v\n", err)
	}
	fmt.Printf("connection established\n")
	defer func() {
		conn.Close()
	}()
	in := bufio.NewReader(conn)

	for {
		digit1, err := in.ReadByte()
		if err == io.EOF || digit1 == '\n' {
			break
		}
		if err != nil {
			log.Printf("unable to read digit: %v", err)
			return
		}
		digit2, err := in.ReadByte()
		if err == io.EOF || digit2 == '\n' {
			break
		}
		if err != nil {
			log.Printf("unable to read digit: %v", err)
			return
		}
		i := ((digit1 - '0') * 10) + (digit2 - '0')
		fmt.Printf(".")
		ch := lut[i]
		_, err = out.WriteString(ch)
		if err != nil {
			log.Printf("error on write: %v\n", err)
			return
		}
		out.Flush()
	}
}

func main() {
	lut = make([]string, 100)
	i := 0
	for ch := 'a'; ch <= 'z'; ch++ {
		lut[i] = string(ch)
		i++
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		lut[i] = string(ch)
		i++
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		lut[i] = string(ch)
		i++
	}
	for ch := '0'; ch <= '9'; ch++ {
		lut[i] = string(ch)
		i++
	}
	for ch := '0'; ch <= '9'; ch++ {
		lut[i] = string(ch)
		i++
	}
	lut[i] = "-"
	i++
	lut[i] = "&"
	i++

	ln, err := net.Listen("tcp", ":31415")
	if err != nil {
		log.Fatalf("unable to listen: %v", err)
	}
	for {
		fmt.Println("listening on port 31415")
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("unable to accept: %v", err)
		}
		out := bufio.NewWriter(conn)
		reader(out)
		conn.Close()
		fmt.Println("connection closed")
	}
}
