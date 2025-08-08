package main

import "net"

func main() {
	lis, err := net.Listen("tcp", ":50052")
}