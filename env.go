package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// SaveToEnv allows overwriting environment variables programmatically in the .env file.
func SaveToEnv(key, value string) error {
	file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the current content of the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Update or add the key-value pair
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
			found = true
			break
		}
	}
	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	// Write the content back to the file
	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	return nil
}
