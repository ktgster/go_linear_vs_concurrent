package linear

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FibonacciInput defines the expected JSON input format.
type FibonacciInput struct {
	Numbers []int `json:"numbers"`
}

// FibonacciOutput defines the JSON output format.
type FibonacciOutput struct {
	Results       []int  `json:"results"`
	ExecutionTime string `json:"execution_time,omitempty"`
}

// fib calculates the nth Fibonacci number recursively.
// Note: This naive implementation is for demonstration purposes.
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// LinearFibHandler accepts a JSON payload with an array of integers,
// sequentially computes the Fibonacci number for each element, and returns the computed
// results in the same order as received along with the total execution time.
func LinearFibHandler(c *gin.Context) {
	// Bind the incoming JSON payload to the FibonacciInput struct.
	var input FibonacciInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// Return an error if the input doesn't match the expected format.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an output slice with the same length as the input array.
	n := len(input.Numbers)
	output := make([]int, n)

	// Record the start time.
	start := time.Now()

	// Process each element sequentially.
	for i, num := range input.Numbers {
		// Compute the Fibonacci number.
		result := fib(num)
		// Write the result to the output slice at the correct index.
		output[i] = result
		// Log the computation for debugging.
		fmt.Printf("Index %d, Value %d, Fibonacci: %d\n", i, num, result)
	}

	// Calculate the total elapsed time.
	elapsed := time.Since(start)
	fmt.Printf("Linear Fibonacci execution took: %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, FibonacciOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
