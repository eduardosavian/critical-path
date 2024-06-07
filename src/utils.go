package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read CSV function
func readCSV(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Split(line, ";")
		matrix = append(matrix, fields)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(matrix) == 0 {
		log.Fatal("The CSV file is empty")
	}

	return matrix
}

// Convert string to int function
func convertToInt(array []string) []int {
	var converted []int

	for _, n := range array {
		n, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			panic("Error converting element:" + err.Error())
		}
		converted = append(converted, n)
	}

	return converted
}

// Format Matrix function from string of the CSV file
func formatMatrix(data []int, size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	dataIndex := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = -1
			} else if i < j {
				matrix[i][j] = data[dataIndex]
				dataIndex++
			} else {
				matrix[i][j] = matrix[j][i]
			}
		}
	}

	return matrix
}

// Print Matrix function
func printMatrix(matrix [][]int) {
	fmt.Println("Matrix")
	for _, array := range matrix {
		for j := range array {
			fmt.Print(array[j], " ")
		}
		fmt.Println()
	}
}
// Calculate the total distance of a route, including the return to the starting city
func calculateTotalDistance(matrix [][]int, route []int) int {
	totalDistance := 0
	numCities := len(route)

	for i := 0; i < numCities-1; i++ {
		totalDistance += matrix[route[i]][route[i+1]]
	}
	// Include the return to the starting city
	totalDistance += matrix[route[numCities-1]][route[0]]

	return totalDistance
}

// Find the best route in the population
func findBestRoute(matrix [][]int, population [][]int) ([]int, int) {
	bestRoute := make([]int, len(population[0]))
	bestDistance := int(^uint(0) >> 1)

	for _, route := range population {
		distance := calculateTotalDistance(matrix, route)
		if distance < bestDistance {
			bestDistance = distance
			copy(bestRoute, route)
		}
	}

	return bestRoute, bestDistance
}
