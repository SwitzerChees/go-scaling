package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
)

var (
	escalateFlag int32
)

func main() {
	http.HandleFunc("/compute/", computeHandler)
	http.HandleFunc("/escalate", escalateHandler)
	http.HandleFunc("/stopescalate", stopEscalateHandler)

	fmt.Println("Server is starting...")
	http.ListenAndServe(":8080", nil)
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the Fibonacci number from the URL path
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fibNum, err := strconv.Atoi(path[2])
	if err != nil {
		http.Error(w, "Invalid Fibonacci number", http.StatusBadRequest)
		return
	}

	// Perform CPU-intensive computations with the specified Fibonacci number
	result := fibonacci(fibNum)

	// Return the fibonacci result
	w.Write([]byte(fmt.Sprintf("Fibonacci(%d) = %d", fibNum, result)))
}

func doIntensiveComputation(n int) {
	_ = fibonacci(n)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func escalateHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		atomic.StoreInt32(&escalateFlag, 1)
		for atomic.LoadInt32(&escalateFlag) == 1 {
			// CPU-intensive loop
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func stopEscalateHandler(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&escalateFlag, 0)
	w.WriteHeader(http.StatusOK)
}
