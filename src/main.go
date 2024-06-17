package main

func main() {
	inputPath := "data/data.csv"
	activities := readActivities(inputPath)
	printInitialData(activities)
	nodes := buildGraph(activities)
	calculateTimes(nodes)
	printResults(nodes)
}
