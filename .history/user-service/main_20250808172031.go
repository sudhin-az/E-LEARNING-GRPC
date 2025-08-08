package main

import (
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "5000")
	if err != nil {
		log.Fatal(err)
	}
}