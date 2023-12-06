package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//make connection
	conn, err := net.Dial("tcp", "localhost:9000")
	if err!= nil {
		log.Fatal(err)
	}

	done:= make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //copy from conn to console
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //to wait for background goroutine by waiting for done
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}