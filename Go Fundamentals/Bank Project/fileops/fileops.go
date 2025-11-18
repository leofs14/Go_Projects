package fileops  // When working with different packages each package needs its own folder with the package name

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) { // When exporting a function to other package it needs to start with an uppercase
	data, err := os.ReadFile(fileName) // _ In Go means a placeholder

	if err != nil { // Nil is when ou have no errors.
		return 1000, errors.New("failed to find balance file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) { //When exporting a function it needs to start with an uppercase
	valueText := fmt.Sprint("Account Balance: ", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}