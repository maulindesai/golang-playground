package main

import (
	"os"
	"testing"
)

func TestDup3(t *testing.T) {
	// Example input data for testing
	content := "line1\nline2\nline1\nline3\nline2\n"
	tempFileName := "temp_test_file.txt"

	// Create a temporary file with sample content
	err := os.WriteFile(tempFileName, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Unexpected error when creating temp file: %v", err)
	}
	// Ensure file is removed after test
	defer os.Remove(tempFileName)

	// Open the file to get the *os.File instance
	file, err := os.Open(tempFileName)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Call the function from dup2.go that processes duplicates
	duplicates := make(map[string]int)
	countLines(file, duplicates)

	// Verify the results, assuming `dup2` returns a map of line frequencies
	expected := map[string]int{
		"line1": 2,
		"line2": 2,
		"line3": 1,
	}
	for key, val := range expected {
		if duplicates[key] != val {
			t.Errorf("For line %q, expected count %d but got %d", key, val, duplicates[key])
		}
	}
}
