package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	//source: https://github.com/dwyl/english-words
	data, err := ioutil.ReadFile("words.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	//fmt.Print(string(data))

	conn, err := net.Dial("tcp", "127.0.0.1:31415")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()

	buff := make([]byte, 1024)
	for {
		n, _ := conn.Read(buff)
		log.Printf("Receive: %s", buff[n:])
	}
}
