package main

func main() {
	activities := readActivities()
	printInitialData(activities)
	nodes := buildGraph(activities)
	calculateTimes(nodes)
	printResults(nodes)
}
