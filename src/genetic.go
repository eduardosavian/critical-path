package main

import (
	"math/rand"
	"time"
	"fmt"
)

// Genetic algorithm function to calculate the best route and distance
func geneticAlgorithm(matrix [][]int, numGenerations int, populationSize int) ([]int, int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	numCities := len(matrix)
	population := initializePopulation(numCities, populationSize)
	bestRoute := make([]int, numCities)
	bestDistance := int(^uint(0) >> 1)

	for generation := 0; generation < numGenerations; generation++ {
		fmt.Print()
		population = evaluateAndSelect(matrix, population)
		population = crossoverAndMutate(population)
		bestInGeneration, bestDistInGeneration := findBestRoute(matrix, population)

		if bestDistInGeneration < bestDistance {
			bestDistance = bestDistInGeneration
			copy(bestRoute, bestInGeneration)
		}
	}

	return bestRoute, bestDistance
}

// Initialize the population with random routes
func initializePopulation(numCities, populationSize int) [][]int {
	population := make([][]int, populationSize)
	for i := 0; i < populationSize; i++ {
		route := rand.Perm(numCities)
		population[i] = route
	}
	return population
}

// Evaluate the fitness of each route and select the best routes
func evaluateAndSelect(matrix [][]int, population [][]int) [][]int {
	populationSize := len(population)
	fitness := make([]int, populationSize)

	for i, route := range population {
		fitness[i] = calculateTotalDistance(matrix, route)
	}

	selectedPopulation := make([][]int, populationSize/2)
	for i := 0; i < populationSize/2; i++ {
		bestIdx := 0
		for j := 1; j < populationSize; j++ {
			if fitness[j] < fitness[bestIdx] {
				bestIdx = j
			}
		}
		selectedPopulation[i] = population[bestIdx]
		fitness[bestIdx] = int(^uint(0) >> 1)
	}

	return selectedPopulation
}

// Perform crossover and mutation to generate new routes
func crossoverAndMutate(population [][]int) [][]int {
	populationSize := len(population)
	newPopulation := make([][]int, populationSize*2)

	for i := 0; i < populationSize; i++ {
		parent1 := population[rand.Intn(populationSize)]
		parent2 := population[rand.Intn(populationSize)]
		child := crossover(parent1, parent2)
		mutate(child)
		newPopulation[i] = parent1
		newPopulation[populationSize+i] = child
	}

	return newPopulation
}

// Crossover two parent routes to create a new child route
func crossover(parent1, parent2 []int) []int {
	numCities := len(parent1)
	child := make([]int, numCities)
	copy(child, parent1)
	start, end := rand.Intn(numCities), rand.Intn(numCities)
	if start > end {
		start, end = end, start
	}

	childPart := make(map[int]bool)
	for i := start; i <= end; i++ {
		childPart[child[i]] = true
	}

	idx := 0
	for i := 0; i < numCities; i++ {
		if !childPart[parent2[i]] {
			for idx >= start && idx <= end {
				idx++
			}
			child[idx] = parent2[i]
			idx++
		}
	}

	return child
}

// Mutate a route by swapping two cities
func mutate(route []int) {
	numCities := len(route)
	if rand.Float64() < 0.4 {
		i, j := rand.Intn(numCities), rand.Intn(numCities)
		route[i], route[j] = route[j], route[i]
	}
}
