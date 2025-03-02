package concurrent

import (
	"fmt"
	"net/http"
	"sync"
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

// fib computes the nth Fibonacci number recursively.
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

// ConcurrentFibHandler accepts a JSON payload with an array of integers,
// concurrently computes the Fibonacci number for each integer,
// and returns the computed results in the same order as received.
func ConcurrentFibHandler(c *gin.Context) {
	// Bind the incoming JSON payload to FibonacciInput.
	var input FibonacciInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// Return a 400 error if JSON binding fails.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Determine the number of input elements.
	n := len(input.Numbers)
	// Create an output slice to store results in the correct order.
	output := make([]int, n)

	// Record the start time.
	start := time.Now()

	// Create a WaitGroup to wait for all goroutines.
	var wg sync.WaitGroup
	wg.Add(n)

	// Process each number concurrently.
	for i, number := range input.Numbers {
		// Capture loop variables.
		go func(idx, num int) {
			defer wg.Done()
			// Compute the Fibonacci number.
			result := fib(num)
			// Save the result in the output slice at the corresponding index.
			output[idx] = result
			// Log the computation for debugging.
			fmt.Printf("Index %d, Value %d, Fibonacci: %d\n", idx, num, result)
		}(i, number)
	}

	// Wait for all computations to complete.
	wg.Wait()

	// Calculate the total execution time.
	elapsed := time.Since(start)
	fmt.Printf("Concurrent Fibonacci execution took %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, FibonacciOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
