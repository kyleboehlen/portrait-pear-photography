package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
	startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println("Generating 120 hashes (this will take a while)...")
	fmt.Printf("Password length: %d characters ✓\n", len(password))

	for i := 0; i < 120; i++ { // 10 years * 12 months
		currentDate := startDate.AddDate(0, i, 0)
		yearMonth := currentDate.Format("2006-01")

		// Generate bcrypt hash
		hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			panic(err)
		}

		_, _ = fmt.Fprintf(f, "\t\"%s\": \"%s\",\n", yearMonth, string(hash))

		// Progress indicator
		if (i+1)%12 == 0 {
			fmt.Printf("Progress: %d/%d years complete\n", (i+1)/12, 10)
		}
	}

	_, _ = fmt.Fprintln(f, "}")
	_, _ = fmt.Fprintln(f, "")
	_, _ = fmt.Fprintf(f, "// Generated on %s\n", time.Now().Format("2006-01-02"))
	_, _ = fmt.Fprintln(f, "// Valid through December 2034. Future me's problem after that.")

	fmt.Printf("\n✓ Successfully generated hashes.go at: %s\n", outputFile)
}
