package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	coursepb "course-service/proto"
	userpb "user-service/proto"
)

func main() {
	r := gin.Default()

	// gRPC Connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInSecure())
	userClient := userpb.NewUserServiceClient(userConn)

	courseConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	courseClient := coursepb.NewCourseServiceClient(courseConn)

	//Routes
	r.POST("/users", func(c *gin.))
}
