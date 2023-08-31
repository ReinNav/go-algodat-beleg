package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Node struct {
	ID                string
	adjacencyList     map[string]*Edge
	distanceFromStart float64
	explored          bool
}

type Edge struct {
	id              string
	sourceNode      *Node
	destinationNode *Node
	weight          float64
}

type Graph struct {
	nodes map[string]*Node
	edges map[string]*Edge
}

func (g *Graph) AddVertex(nodeId string) {
	// Check if node exists
	if _, exists := g.nodes[nodeId]; !exists {
		g.nodes[nodeId] = &Node{
			ID:            nodeId,
			adjacencyList: make(map[string]*Edge),
		}
	}
}

func NewGraph() Graph {
	return Graph{
		nodes: make(map[string]*Node),
		edges: make(map[string]*Edge),
	}
}

func (g *Graph) NumVertices() int {
	return len(g.nodes)
}

func (g *Graph) NumEdges() int {
	return len(g.edges)
}

func (g *Graph) BFS(s string) map[string]int {
	result := make(map[string]int)

	for _, node := range g.nodes {
		node.explored = false
	}

	if _, ok := g.nodes[s]; !ok {
		return result
	}

	g.nodes[s].explored = true
	result[s] = 0

	queue := NewQueue()
	queue.Enqueue(g.nodes[s])

	for !queue.IsEmpty() {
		v := queue.Dequeue()

		for _, edge := range v.adjacencyList {
			w := edge.destinationNode

			if !w.explored {
				w.explored = true
				result[w.ID] = result[v.ID] + 1
				queue.Enqueue(w)
			}
		}
	}

	return result
}

func (g *Graph) DFS(s string) map[string]bool {
	visited := make(map[string]bool)

	// Start node
	startNode, ok := g.nodes[s]
	if !ok {
		return nil
	}

	stack := NewStack()
	stack.Push(startNode)

	for !stack.IsEmpty() {
		// Pop
		currentNode := stack.Pop()

		if currentNode.explored {
			continue
		}

		// Matk as explored & reachable
		currentNode.explored = true
		visited[currentNode.ID] = true

		// add unexplored nodes in adjencency list of currentNOde to the stack
		for _, edge := range currentNode.adjacencyList {
			if !edge.destinationNode.explored {
				stack.Push(edge.destinationNode)
			}
		}
	}

	result := make(map[string]bool)

	for _, node := range g.nodes {
		if _, ok := visited[node.ID]; ok {
			result[node.ID] = true
		} else {
			result[node.ID] = false
		}
	}

	return result
}

// Djikstra: Find shortest path to each node from start node startNodeID
func (g *Graph) Dijkstra(s string) map[string]float64 {
	distances := make(map[string]float64)

	for id := range g.nodes {
		if id == s {
			distances[id] = 0
		} else {
			distances[id] = math.Inf(1)
		}
		g.nodes[id].distanceFromStart = distances[id]
	}

	minHeap := NewMinHeap()
	minHeap.Insert(g.nodes[s])

	for !minHeap.IsEmpty() {
		// Extract min node
		u := minHeap.ExtractMin()

		// Update distances for nodes adjacent to u
		for _, edge := range u.adjacencyList {
			v := edge.destinationNode

			// Compare, if current path smaller, replace
			if u.distanceFromStart+edge.weight < v.distanceFromStart {
				v.distanceFromStart = u.distanceFromStart + edge.weight // NEw shortest
				distances[v.ID] = v.distanceFromStart
				minHeap.Insert(v)
			}
		}
	}

	return distances
}

func (g *Graph) AddDirectedEdge(fromNodeId string, toNodeId string, weight float64) {
	fromNode := g.nodes[fromNodeId]
	toNode := g.nodes[toNodeId]

	if fromNode == nil || toNode == nil {
		fmt.Printf("The edge (%v -> %v) is invalid\n", fromNodeId, toNodeId)
		return
	}

	// Check if edge already exist
	for _, edge := range fromNode.adjacencyList {
		if edge.destinationNode == toNode {
			fmt.Printf("The edge (%v -> %v) already exists.\n", fromNodeId, toNodeId)
			return
		}
	}

	// Create new edge
	newEdge := &Edge{
		id:              fromNode.ID + "," + toNode.ID,
		sourceNode:      fromNode,
		destinationNode: toNode,
		weight:          weight,
	}
	g.edges[newEdge.id] = newEdge

	fromNode.adjacencyList[toNodeId] = newEdge
}

func (g *Graph) AddUndirectedEdge(nodeId1, nodeId2 string, length float64) {

	// CHeck if nodes exist
	node1, node1Exists := g.nodes[nodeId1]
	node2, node2Exists := g.nodes[nodeId2]

	if !node1Exists || !node2Exists {
		fmt.Printf("One or both nodes do not exist in graph\n")
		return
	}

	edgeIdToNode2 := nodeId1 + "," + nodeId2
	edgeToNode2 := &Edge{
		id:              edgeIdToNode2,
		sourceNode:      node1,
		destinationNode: node2,
		weight:          length,
	}

	edgeIdToNode1 := nodeId2 + "," + nodeId1
	edgeToNode1 := &Edge{
		id:              edgeIdToNode1,
		sourceNode:      node2,
		destinationNode: node1,
		weight:          length,
	}

	_, edgeToNode2Exists := g.edges[edgeIdToNode2]
	_, edgeToNode1Exists := g.edges[edgeIdToNode1]

	if edgeToNode2Exists || edgeToNode1Exists {
		fmt.Printf("One or both edges already exist in graph\n")
		return
	}

	node1.adjacencyList[nodeId2] = edgeToNode2
	node2.adjacencyList[nodeId1] = edgeToNode1

	// Add edge to list of edfes
	g.edges[edgeIdToNode1] = edgeToNode1
	g.edges[edgeIdToNode2] = edgeToNode2
}

func (g *Graph) UCC() map[string]int {
	visited := make(map[string]bool)
	result := make(map[string]int)
	componentID := 0

	// mark all nodes unvisited
	for nodeId := range g.nodes {
		visited[nodeId] = false
	}

	for nodeId, node := range g.nodes {
		if !visited[nodeId] {
			stack := NewStack()
			stack.Push(node)

			for !stack.IsEmpty() {
				current := stack.Pop()

				visited[current.ID] = true
				// assign it to current cc
				result[current.ID] = componentID

				// check adjecency
				for _, edge := range current.adjacencyList {
					if !visited[edge.destinationNode.ID] {
						stack.Push(edge.destinationNode)
					}
				}
			}

			// For start node component id = 0
			componentID++
		}
	}

	return result
}

func (g *Graph) TopoSort() map[string]int {
	stack := NewStack()
	visited := make(map[string]bool)

	for _, node := range g.nodes {
		if !visited[node.ID] {
			g.DFSTopo(node, visited, stack)
		}
	}

	order := make(map[string]int)
	i := 1
	for !stack.IsEmpty() {
		node := stack.Pop()
		order[node.ID] = i
		i++
	}
	return order
}

func (g *Graph) DFSTopo(node *Node, visited map[string]bool, stack *Stack) {
	visited[node.ID] = true
	for _, edge := range node.adjacencyList {
		if !visited[edge.destinationNode.ID] {
			g.DFSTopo(edge.destinationNode, visited, stack)
		}
	}
	stack.Push(node)
}

func initGraph9(filename string) Graph {

	graph := NewGraph()

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",") // f[0]:id2 ,,
			var length float64         //edgeLength
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l //edgeLength{float64: l}
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddDirectedEdge(id1, f[0], length)
		}

	}
	return graph
}

func initWebgraph(t *testing.T) Graph {
	// for a sanity check:
	//     count on the command line the number of edges and vertices by
	// grep -E -v "^#" ~/Downloads/web-Google.txt | wc -l
	// grep -E -v "^#" ~/Downloads/web-Google.txt | sed -E 's/([[:digit:]]+)[[:space:]]+([[:digit:]]+).*/\1\n\2/' | sort | uniq | wc -l

	graph := NewGraph()

	file, _ := os.Open("res/web-Google.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	start := time.Now()
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if !strings.HasPrefix((fields[0][0:1]), "#") && len(fields) == 2 {
			graph.AddVertex(fields[0])
			graph.AddVertex(fields[1])
			graph.AddDirectedEdge(fields[0], fields[1], 1.)
		}
		if (i % 1000000) == 0 {
			elapsed := time.Since(start)
			t.Logf("last took %s\n", elapsed)
			t.Logf("progess: %v\n", i)
			start = time.Now()
		}
		i++
	}
	t.Logf("%v lines processed\n", i)

	return graph
}

func main() {
	// Graph Problem 9.8
	fmt.Println("\n\nGraph: problem9.8.txt")
	graph98 := initGraph9("res/problem9.8.txt")

	fmt.Printf("Number of nodes: %v\n", graph98.NumVertices())
	fmt.Printf("Number of edges: %v\n", graph98.NumEdges())
	//fmt.Printf("Topo-Sort: %v\n", graph98.TopoSort())

	djikstraResultFrom1 := graph98.Dijkstra("1")

	keys := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}
	for _, key := range keys {
		if val, ok := djikstraResultFrom1[key]; ok {
			fmt.Printf("Shortest path from node 1 to node %s: %.3f\n", key, val)
		}
	}

}
