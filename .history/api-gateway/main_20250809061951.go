package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sudhin-az/E-LEARNING-GRPC/course-service/proto"
	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	// gRPC Connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := proto.NewCourseServiceClient(userConn)

	courseConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	courseClient := proto.NewCourseServiceClient(courseConn)

	//Routes
	r.POST("/users", func(c *gin.Context) {
		var req proto.UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := userClient.c(context.Background(), &req)
		c.JSON(http.StatusOK, res)
	})
	r.POST("/courses", func(c *gin.Context) {
		var req proto.CourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := courseClient.CreateCourse(context.Background(), &req)
		c.JSON(http.StatusOK, res)
	})

	r.POST("/enroll", func(c *gin.Context) {
		var req proto.EnrollRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := courseClient.EnrollCourse(context.Background(), &req)
		c.JSON(http.StatusOK, res)
	})
	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
