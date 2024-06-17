package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

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
