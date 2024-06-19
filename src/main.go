package main

func main() {
	inputPath := "data/data.csv"
	activities, err := readActivities(inputPath)
	if err != nil {
		panic(err)
	}
	printInitialData(activities)
	nodes := buildGraph(activities)
	calculateTimes(nodes)
	printResults(nodes)
}
