package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func main() {

	fmt.Println("Hello, playground")
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// init z to unit
	//z := new(Lft)
	//z.q.SetInt64(1)
	//z.t.SetInt64(1)
	z, err := rFromWfile("q.txt", "r.txt", "s.txt", "t.txt")
	if err != nil {
		fmt.Println(err)
	}

	// lfts generator
	var k int64
	lfts := func() *Lft {
		k++
		r := new(Lft)
		r.q.SetInt64(k)
		r.r.SetInt64(4*k + 2)
		r.t.SetInt64(2*k + 1)
		return r
	}
	// boot up tcp server
	listener, err := net.Listen("tcp", ":3141")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	// accept new connection
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("tcp server accept error", err)
	}

	// stream
	for {
		select {
		case <-done:
			fmt.Printf("exiting size %+v", unsafe.Sizeof(z))
			wToFile(z)
			conn.Close()
			os.Exit(0)
		default:
			y := z.Next()
			if z.Safe(y) {
				fmt.Print(y)
				conn.Write([]byte(y.Text(10)))
				z = z.Prod(y)
			} else {
				z = z.Comp(lfts())
			}
		}
	}
}

func wToFile(pi *Lft) error {
	err := ioutil.WriteFile("q.txt", []byte(pi.q.Text(10)), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("r.txt", []byte(pi.r.Text(10)), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("s.txt", []byte(pi.s.Text(10)), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("t.txt", []byte(pi.t.Text(10)), 0644)
	if err != nil {
		return err
	}
	return err
}

func rFromWfile(fq, fr, fs, ft string) (*Lft, error) {
	z := new(Lft)
	z.q.SetInt64(1)
	z.t.SetInt64(1)
	dq, err := ioutil.ReadFile(fq)
	if err != nil {
		return z, err
	}
	z.q.SetString(string(dq), 10)
	dr, err := ioutil.ReadFile(fr)
	if err != nil {
		return z, err
	}
	z.r.SetString(string(dr), 10)
	ds, err := ioutil.ReadFile(fs)
	if err != nil {
		return z, err
	}
	z.s.SetString(string(ds), 10)
	dt, err := ioutil.ReadFile(ft)
	if err != nil {
		return z, err
	}
	z.t.SetString(string(dt), 10)
	return z, err
}
