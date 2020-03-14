package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":3141")
	if err != nil {
		log.Fatalf("unable to listen: %v", err)
	}
	for {
		fmt.Println("listening on port 3141")
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("unable to accept: %v", err)
		}
		fmt.Println("connection established")
		out := bufio.NewWriter(conn)
		emit(out)
		fmt.Println("connection closed")
	}
}

func emit(out *bufio.Writer) {
	for {
		i := rand.Intn(9)
		s := fmt.Sprintf("%v", i)
		_, err := out.WriteString(s)
		if err != nil {
			log.Printf("unable to write: %v\n", err)
			return
		}
		time.Sleep(time.Millisecond * 1)
		fmt.Printf(s)
		out.Flush()
	}
}
