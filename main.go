package main

import (
	"yourmodule/handlers/concurrent" // replace "yourmodule" with your actual module name
	"yourmodule/handlers/linear"     // replace "yourmodule" with your actual module name

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Sequential (linear) endpoints.
	router.POST("/linear_factorial", linear.LinearFactorialHandler)
	router.POST("/linear_fib", linear.LinearFibHandler)
	router.POST("/linear_primechecker", linear.LinearPrimeCheckerHandler)

	// Concurrent endpoints.
	router.POST("/concurrent_factorial", concurrent.ConcurrentFactHandler) // assuming this exists
	router.POST("/concurrent_fib", concurrent.ConcurrentFibHandler)        // assuming this exists
	router.POST("/concurrent_primechecker", concurrent.ConcurrentPrimeHandler)

	router.Run(":8080")
}
