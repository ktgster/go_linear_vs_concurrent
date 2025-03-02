# Go: Linear vs Concurrent Programming Benchmark

This project demonstrates and benchmarks two approaches for handling CPU-intensive tasks in Golang: **linear (sequential) looping** versus **concurrent programming**. We compare the performance of both techniques using three computationally intensive operations:

- **Fibonacci Number Calculation**
- **Factorial Calculation**
- **Prime Checking**

## Overview

### Linear (Sequential) Processing
In the linear approach, the program processes each element in the input array one at a time. This means that each computation must finish before the next one begins. While this method is simple and straightforward, it does not take full advantage of multi-core processors, making it slower for CPU-intensive tasks.

### Concurrent Programming
Concurrent programming in Go uses lightweight **goroutines** that allow multiple computations to run simultaneously. In our concurrent implementation, each element in the input array is processed in its own goroutine, and a `sync.WaitGroup` is used to wait for all goroutines to finish before returning the results. This approach can significantly improve performance by parallelizing work across available CPU cores.

## Tested Operations

1. **Fibonacci Number Calculation**  
   - *Linear Approach:* Computes Fibonacci numbers one after the other using a for-loop.
   - *Concurrent Approach:* Uses goroutines to compute Fibonacci numbers concurrently.  
   > **Note:** The naive recursive Fibonacci function is very inefficient for high input values. For example, calculating Fibonacci(50) can be time-consuming, and higher values may cause stack overflows or excessive computation times.

2. **Factorial Calculation**  
   - *Linear Approach:* Computes factorials sequentially, using `math/big` to handle very large numbers.
   - *Concurrent Approach:* Launches goroutines for factorial computation for each input number concurrently.
   
3. **Prime Checking**  
   - *Linear Approach:* Checks each number one after the other to see if it is prime.
   - *Concurrent Approach:* Uses goroutines to check multiple numbers for primality at the same time.

## Performance Comparison

In our tests with the Fibonacci function for input `[50,50,50,50,50,50,50,50,50,50,50]`:
- **Concurrent Fibonacci:** Approximately **1 minute 21 seconds** (≈81.09 seconds)
- **Linear Fibonacci:** Approximately **10 minutes 1 second** (≈601.41 seconds)

This results in roughly a **7.4x speedup** using concurrent programming.  
*Actual improvements depend on the number of CPU cores, the specific workload, and the overhead of goroutine management.*

## How It Works

- **Linear Implementation:**  
  Iterates over an array and processes each element sequentially.

- **Concurrent Implementation:**  
  Launches a goroutine for each element, allowing multiple computations to run in parallel. A `sync.WaitGroup` is used to ensure that the main process waits for all goroutines to complete before returning the final result.

## Getting Started

### Prerequisites
- Go (version 1.16+ recommended)
- [Gin](https://github.com/gin-gonic/gin) framework for creating the HTTP server

### Setup
1. **Clone the repository and navigate to the project directory:**
   ```bash
   git clone <repository-url>
   cd Golang
