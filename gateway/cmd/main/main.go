package main

import (
	"github.com/bhtoan2204/gateway/internal/initialize"
) // gin-swagger middleware
// swagger embed files
// @title Todo Application
// @description This is a todo list management application
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	initialize.Run()
}
