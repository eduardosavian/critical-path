package main

type Activity struct {
	Name       string
	Duration   int
	Precedents []string
}

type Node struct {
	Name       string
	ES, EF     int
	LS, LF     int
	Duration   int
	Precedents []string
	Successors []*Node
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
	for i := len(sortedNodes) - 1; i >= 0; i-- {
		node := sortedNodes[i]
		node.LF = maxEF
		for _, successor := range node.Successors {
			if successor.LS < node.LF {
				node.LF = successor.LS
			}
		}
		node.LS = node.LF - node.Duration
	}
}

func findCriticalPath(nodes map[string]*Node) []string {
	var criticalPath []string
	startNode := findStartNode(nodes)
	var visit func(node *Node)
	visit = func(node *Node) {
		criticalPath = append(criticalPath, node.Name)
		for _, successor := range node.Successors {
			if successor.ES == successor.LS { // Zero slack check
				visit(successor)
				break
			}
		}
	}
	visit(startNode)
	return criticalPath
}

func findStartNode(nodes map[string]*Node) *Node {
	for _, node := range nodes {
		if len(node.Precedents) == 0 {
			return node
		}
	}
	return nil
}
