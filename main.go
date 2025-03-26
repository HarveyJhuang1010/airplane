package main

import (
	cmd "airplane/internal/application/command"
)

// @title						Airplane - Flight Booking System
// @version						1.0.0
// @description					This is a RESTFUL API documentation of Flight Booking System.
// @host						localhost:8080
// @BasePath					/api/v1
// @schemes						http
func main() {
	cmd.Execute()
}
