package main

import (
	_ "github.com/bhtoan2204/gateway/cmd/main/docs" // Import the docs
	"github.com/bhtoan2204/gateway/internal/initialize"
)

// @title           API Gateway Service
// @version         1.0
// @description     This is the API Gateway service for YouTube Clone project
// @contact.name    Banh Hao Toan
// @contact.email   banhhaotoan2002@gmail.com
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	initialize.Run()
}
