package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	userConn, _ := grpc.Dial("localhost:50051", grpc.wis)
}
