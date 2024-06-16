package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Activity struct {
	Name       string
	Duration   int
	Precedents []string
}

type Node struct {
	Name       string
	ES, EF     int // Early Start, Early Finish
	LS, LF     int // Late Start, Late Finish
	Duration   int
	Precedents []string
	Successors []*Node
}

func readActivities() []Activity {
	var activities []Activity
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter activity (name duration precedents), or 'done' to finish:")
		scanner.Scan()
		input := scanner.Text()

		if input == "done" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 2 {
			fmt.Println("Invalid input, please enter again.")
			continue
		}

		name := parts[0]
		var duration int
		if _, err := fmt.Sscanf(parts[1], "%d", &duration); err != nil {
			fmt.Println("Invalid input, please enter again.")
			continue
		}

		precedents := []string{}
		if len(parts) > 2 && parts[2] != "-" {
			precedents = strings.Split(parts[2], ",")
		}

		activities = append(activities, Activity{
			Name:       name,
			Duration:   duration,
			Precedents: precedents,
		})
	}
	return activities
}

func buildGraph(activities []Activity) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, activity := range activities {
		nodes[activity.Name] = &Node{
			Name:       activity.Name,
			Duration:   activity.Duration,
			Precedents: activity.Precedents,
		}
	}
	for _, node := range nodes {
		for _, precedent := range node.Precedents {
			if precedent != "" {
				nodes[precedent].Successors = append(nodes[precedent].Successors, node)
			}
		}
	}
	return nodes
}

func topologicalSort(nodes map[string]*Node) []*Node {
	var sorted []*Node
	visited := make(map[string]bool)

	var visit func(node *Node)
	visit = func(node *Node) {
		if visited[node.Name] {
			return
		}
		visited[node.Name] = true
		for _, successor := range node.Successors {
			visit(successor)
		}
		sorted = append(sorted, node)
	}

	for _, node := range nodes {
		visit(node)
	}

	// Reverse the sorted list for backward pass
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	return sorted
}

func calculateTimes(nodes map[string]*Node) {
	// Topological sorting for proper order of processing
	sortedNodes := topologicalSort(nodes)

	// Forward Pass: Calculate ES and EF
	for _, node := range sortedNodes {
		node.ES = 0
		for _, precedent := range node.Precedents {
			if precedent != "" {
				if nodes[precedent].EF > node.ES {
					node.ES = nodes[precedent].EF
				}
			}
		}
		node.EF = node.ES + node.Duration
	}

	// Find the maximum EF
	maxEF := 0
	for _, node := range nodes {
		if node.EF > maxEF {
			maxEF = node.EF
		}
	}

	// Backward Pass: Calculate LS and LF
	// for _, node := range sortedNodes {
	// 	node.LF = maxEF
	// }

	// for _, node := range sortedNodes {
	// 	node.LS = node.LF - node.Duration
	// 	for _, successor := range node.Successors {
	// 		if successor.LS < node.LF {
	// 			node.LF = successor.LS
	// 		}
	// 	}
	// 	node.LS = node.LF - node.Duration
	// }
}

func findCriticalPath(nodes map[string]*Node) []string {
	var criticalPath []string
	for _, node := range nodes {
		if node.ES == node.LS {
			criticalPath = append(criticalPath, node.Name)
		}
	}
	return criticalPath
}

func printInitialData(activities []Activity) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Activity\tDuration\tPrecedents\t")
	for _, activity := range activities {
		fmt.Fprintf(writer, "%s\t%d\t%s\t\n", activity.Name, activity.Duration, strings.Join(activity.Precedents, ","))
	}
	writer.Flush()
}

func printResults(nodes map[string]*Node) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Activity\tES (Early Start)\tEF (Early Finish)\tLS (Late Start)\tLF (Late Finish)\tSlack\t")
	for _, node := range nodes {
		slack := node.LS - node.ES
		fmt.Fprintf(writer, "%s\t%d\t%d\t%d\t%d\t%d\t\n", node.Name, node.ES, node.EF, node.LS, node.LF, slack)
	}
	writer.Flush()

	criticalPath := findCriticalPath(nodes)
	fmt.Println("\nCritical Path:", strings.Join(criticalPath, " -> "))
}

func main() {
	activities := readActivities()
	printInitialData(activities)
	nodes := buildGraph(activities)
	calculateTimes(nodes)
	printResults(nodes)
}
