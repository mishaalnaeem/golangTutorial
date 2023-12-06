package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
)

//mapping for clients
var client = make(map[string]net.Conn) //string to net.Conn mapping
var messages = make(chan message)

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

func broadcaster() {
	//needs a channel for messages from connections 

	for {
		select {
		case msg := <- messages:
			for _, conn := range clients { //iterate over clients
				//send message if client is not sneder
				if msg.address == conn.RemoteAddr().String() {
					continue
				}
				fmt.fPrintln(conn, msg.text)

			}
		}

		//leaving messages
	}

}

func handler(conn net.Conn) {
	//add connection record, send messages <- handles each client

	clients[conn.RemoteAddr().String()] = conn

	//send a client joining message
	messages <- newMessage("joined", conn)

	input := bufio.NewScanner(conn)
	for input.Scan () {
		messages <- newMessage(": "+input.Text(), conn)
	}

	delete(clients, conn.Remote().String())

	conn.Close()

}

func newMessage(msg string, conn net.Conn) message {
	addr := conn.RemoteAddr.String()
	return {
		text: addr + msg,
		address: addr,
	}
}