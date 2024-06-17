package main

import "fmt"

func main() {
	inputPath := "data/data.csv"
	activities, err := readActivities(inputPath)
	fmt.Print(err)
	printInitialData(activities)
	nodes := buildGraph(activities)
	calculateTimes(nodes)
	printResults(nodes)
}
