package concurrent

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// FactorialInput defines the expected JSON input format.
type FactorialInput struct {
	Numbers []int `json:"numbers"`
}

// FactorialOutput defines the JSON output format.
type FactorialOutput struct {
	Results       []int  `json:"results"`
	ExecutionTime string `json:"execution_time,omitempty"`
}

// fact computes the factorial of n iteratively.
// For n < 0, it returns 0 (you might want to handle negative inputs differently).
func fact(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// ConcurrentFactHandler accepts a JSON payload with an array of integers,
// concurrently computes the factorial for each integer,
// and returns the computed results in the same order as received.
func ConcurrentFactHandler(c *gin.Context) {
	// Bind the incoming JSON payload to FactorialInput.
	var input FactorialInput
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
			// Compute the factorial.
			result := fact(num)
			// Save the result in the output slice at the corresponding index.
			output[idx] = result
			// Log the computation for debugging.
			fmt.Printf("Index %d, Value %d, Factorial: %d\n", idx, num, result)
		}(i, number)
	}

	// Wait for all computations to complete.
	wg.Wait()

	// Calculate the total execution time.
	elapsed := time.Since(start)
	fmt.Printf("Concurrent Factorial execution took %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, FactorialOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
