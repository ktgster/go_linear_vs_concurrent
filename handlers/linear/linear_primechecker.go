package linear

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PrimeCheckerInput defines the expected JSON input format.
type PrimeCheckerInput struct {
	Numbers []int `json:"numbers"`
}

// PrimeCheckerOutput defines the JSON output format.
type PrimeCheckerOutput struct {
	Results       []bool `json:"results"`
	ExecutionTime string `json:"execution_time,omitempty"`
}

// isPrime checks if a given number n is prime.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// LinearPrimeCheckerHandler accepts a JSON payload with an array of integers,
// sequentially checks if each number is prime, and returns the results in the same order as received,
// along with the total execution time.
func LinearPrimeCheckerHandler(c *gin.Context) {
	// Bind the incoming JSON payload to the PrimeCheckerInput struct.
	var input PrimeCheckerInput
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

	// Process each element sequentially.
	for i, num := range input.Numbers {
		// Check if the number is prime.
		result := isPrime(num)
		// Write the result to the output slice at the correct index.
		output[i] = result
		// Log the computation for debugging.
		fmt.Printf("Index %d, Value %d, isPrime: %t\n", i, num, result)
	}

	// Calculate the total elapsed time.
	elapsed := time.Since(start)
	fmt.Printf("Linear prime checking execution took: %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, PrimeCheckerOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
