package concurrent

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// PrimeInput defines the expected JSON input format.
type PrimeInput struct {
	Numbers []int `json:"numbers"`
}

// PrimeOutput defines the JSON output format.
type PrimeOutput struct {
	Results       []bool `json:"results"`
	ExecutionTime string `json:"execution_time,omitempty"`
}

// isPrime checks if a given number n is prime.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// Check divisibility from 2 up to sqrt(n)
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ConcurrentPrimeHandler accepts a JSON payload with an array of integers,
// concurrently checks whether each number is prime,
// and returns an array of booleans in the same order as received,
// along with the total execution time.
func ConcurrentPrimeHandler(c *gin.Context) {
	// Bind the incoming JSON payload to the PrimeInput struct.
	var input PrimeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// Return an error if the input doesn't match the expected format.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an output slice with the same length as the input array.
	n := len(input.Numbers)
	output := make([]bool, n)

	// Record the start time.
	start := time.Now()

	// Use a WaitGroup to synchronize the concurrent computations.
	var wg sync.WaitGroup
	wg.Add(n)

	// Process each element concurrently.
	for i, num := range input.Numbers {
		// Capture the loop variables for the goroutine.
		go func(idx, number int) {
			defer wg.Done()
			// Check if the number is prime.
			result := isPrime(number)
			// Write the result to the output slice at the correct index.
			output[idx] = result
			// Log the computation for debugging.
			fmt.Printf("Index %d, Value %d, IsPrime: %v\n", idx, number, result)
		}(i, num)
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// Calculate the total elapsed time.
	elapsed := time.Since(start)
	fmt.Printf("Concurrent Prime check execution took: %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, PrimeOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
