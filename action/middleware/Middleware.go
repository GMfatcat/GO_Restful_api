package middleware

import (
	"fmt"
	"os"
	"time"
)

var logFileName string = "log.txt"

/* Logger for all Client Query and Backend Response */

func QueryLogger(user, query string) error {

	currentTime := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("Time: %s, User: %s, Query: %s", currentTime, user, query)
	// Print and save to log
	fmt.Println(logEntry)
	err := writeToLogFile(logEntry)
	if err != nil {
		return err
	}

	return nil
}

// HOW TO CONNECT TO RESPONSE TBD
func ResponseLogger(query string, httpCode int) error {

	currentTime := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("Time: %s, Query: %s, HTTP Code: %d", currentTime, query, httpCode)
	// Print and save to log
	fmt.Println(logEntry)
	err := writeToLogFile(logEntry)
	if err != nil {
		return err
	}

	return nil
}

func writeToLogFile(entry string) error {
	// Open log file, create it if it doesn't exists
	// 0644 means non-owner can only read it
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write entry to log
	_, err = file.WriteString(entry + "\n")
	if err != nil {
		return err
	}

	return nil
}
