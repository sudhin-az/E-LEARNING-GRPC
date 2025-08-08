package main

import (
	"context"
	coursepb "course-service/proto"
	"log"
	"net/http"
	userpb "user-service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	// gRPC Connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInSecure())
	userClient := userpb.NewUserServiceClient(userConn)

	courseConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	courseClient := coursepb.NewCourseServiceClient(courseConn)

	//Routes
	r.POST("/users", func(c *gin.Context) {
		var req userpb.UserRequest
		if err := c.ShouldBindJson(&req); err != nil {
			c.Json(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := userClient.CreateUser(context.Background(), &req)
		c.JSON(http.StatusOK, res)
	})
	r.POST("/courses", func(c *gin.Context) {
		var req coursepb.CourseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := courseClient.CreateCourse(context.Background(), &req)
		c.JSON(http.StatusOk, res)
	})

	r.POST("/enroll", func(c *gin.Context) {
		var req coursepb.EnrollRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := courseClient.EnrollCourse(context.Background(), &req)
		c.JSON(http.StatusOk, res)
	})
	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
