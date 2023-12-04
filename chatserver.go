package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
)

func main() {
	//open port to listen to
	listen, err := net.Listen("tcp", "localhost:9000")

	//error handling of port opening
	if err != nil {
		log.Fatal(err)
	}

	//run a go routine which handles messages incoming and outgoing
	//a broadcaster
	go broadcaster()

	//listen for connections
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		//go routine for handling each connection
		go handler(conn)
	}
}

func broadcaster(){
	//needs a channel for messages from connections 
}

func handler(conn net,Conn){
	//add connection record
}