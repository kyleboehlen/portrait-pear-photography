package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type hashResult struct {
	yearMonth string
	hash      string
	index     int
}

func main() {
	// Check if password was provided
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: go run main.go <password>")
		_, _ = fmt.Fprintln(os.Stderr, "Example: go run main.go 'your-64-char-password'")
		os.Exit(1)
	}

	password := os.Args[1]

	// Validate password length
	if len(password) < 64 {
		_, _ = fmt.Fprintf(os.Stderr, "Error: Password must be at least 64 characters long (got %d)\n", len(password))
		_, _ = fmt.Fprintln(os.Stderr, "This isn't just security theater - we're going for the heat death of the universe flex here.")
		os.Exit(1)
	}

	// Bcrypt cost
	cost := 17

	// Get directory of the command that called this script
	execPath, _ := os.Getwd()
	scriptDir := filepath.Dir(execPath)

	// Output file in the same directory
	outputFile := filepath.Join(scriptDir, "backend/services/auth/hashes.go")

	// Create/truncate the output file
	f, err := os.Create(outputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// Write package and map declaration
	_, _ = fmt.Fprintln(f, "package auth")
	_, _ = fmt.Fprintln(f, "")
	_, _ = fmt.Fprintln(f, "// Bcrypt hashes for admin password, rotated monthly because I can.")
	_, _ = fmt.Fprintln(f, "// Format: YYYY-MM")
	_, _ = fmt.Fprintln(f, "// 420 bits of entropy + bcrypt cost of 17. You do the math, I'll take those odds.")
	_, _ = fmt.Fprintln(f, "var theForbiddenHashArray = map[string]string{")

	// Generate hashes for 10 years, monthly
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	fmt.Println("Generating 120 hashes concurrently (this will still take a while, but less while)...")
	fmt.Printf("Password length: %d characters ✓\n", len(password))
	fmt.Printf("Using %d CPU cores\n", runtime.NumCPU())

	// Set up concurrency
	numWorkers := runtime.NumCPU()
	jobs := make(chan int, 120)
	results := make(chan hashResult, 120)
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				currentDate := startDate.AddDate(0, i, 0)
				yearMonth := currentDate.Format("2006-01")

				// Generate bcrypt hash with month appended
				hash, err := bcrypt.GenerateFromPassword([]byte(password+yearMonth), cost)
				if err != nil {
					panic(err)
				}

				results <- hashResult{
					yearMonth: yearMonth,
					hash:      string(hash),
					index:     i,
				}

				// Progress indicator
				if (i+1)%12 == 0 {
					fmt.Printf("Progress: %d/%d years complete\n", (i+1)/12, 10)
				}
			}
		}()
	}

	// Send jobs
	go func() {
		for i := 0; i < 120; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Gather all results
	var hashResults []hashResult
	for result := range results {
		hashResults = append(hashResults, result)
	}

	// Sort results by index to maintain chronological order
	sort.Slice(hashResults, func(i, j int) bool {
		return hashResults[i].index < hashResults[j].index
	})

	// Write sorted results to file
	for _, result := range hashResults {
		_, _ = fmt.Fprintf(f, "\t\"%s\": \"%s\",\n", result.yearMonth, result.hash)
	}

	_, _ = fmt.Fprintln(f, "}")
	_, _ = fmt.Fprintln(f, "")
	_, _ = fmt.Fprintf(f, "// Generated on %s\n", time.Now().Format("2006-01-02"))
	_, _ = fmt.Fprintln(f, "// Valid through December 2034. Future me's problem after that.")

	fmt.Printf("\n✓ Successfully generated hashes.go at: %s\n", outputFile)
}
