package linear

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FactorialInput defines the expected JSON input format.
type FactorialInput struct {
	Numbers []int `json:"numbers"`
}

// FactorialOutput defines the JSON output format.
type FactorialOutput struct {
	Results       []string `json:"results"` // Using string to accommodate big numbers.
	ExecutionTime string   `json:"execution_time,omitempty"`
}

// fact computes the factorial of n using big.Int for arbitrary precision.
func fact(n int) *big.Int {
	result := big.NewInt(1)
	if n < 2 {
		return result
	}
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// LinearFactorialHandler accepts a JSON payload with an array of integers,
// sequentially computes the factorial for each element, and returns the computed
// results in the same order as received along with the total execution time.
func LinearFactorialHandler(c *gin.Context) {
	// Bind the incoming JSON payload to the FactorialInput struct.
	var input FactorialInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// Return an error if the input doesn't match the expected format.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an output slice with the same length as the input array.
	n := len(input.Numbers)
	output := make([]string, n)

	// Record the start time.
	start := time.Now()

	// Process each element sequentially.
	for i, num := range input.Numbers {
		// Compute the factorial.
		result := fact(num)
		// Write the result to the output slice at the correct index.
		output[i] = result.String()
		// Log the computation for debugging.
		fmt.Printf("Index %d, Value %d, Factorial: %s\n", i, num, result.String())
	}

	// Calculate the total elapsed time.
	elapsed := time.Since(start)
	fmt.Printf("Linear Factorial execution took: %s\n", elapsed)

	// Return the computed results and execution time as JSON.
	c.JSON(http.StatusOK, FactorialOutput{
		Results:       output,
		ExecutionTime: elapsed.String(),
	})
}
