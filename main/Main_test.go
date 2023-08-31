package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_problem98test(t *testing.T) {
	graph := NewGraph()

	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddVertex("7")
	graph.AddVertex("8")

	graph.AddUndirectedEdge("1", "2", 1)
	graph.AddUndirectedEdge("1", "8", 2)
	graph.AddUndirectedEdge("2", "3", 1)
	graph.AddUndirectedEdge("3", "2", 1)
	graph.AddUndirectedEdge("3", "4", 1)
	graph.AddUndirectedEdge("4", "3", 1)
	graph.AddUndirectedEdge("4", "5", 1)
	graph.AddUndirectedEdge("5", "4", 1)
	graph.AddUndirectedEdge("5", "6", 1)
	graph.AddUndirectedEdge("6", "5", 1)
	graph.AddUndirectedEdge("6", "7", 1)
	graph.AddUndirectedEdge("7", "6", 1)
	graph.AddUndirectedEdge("7", "8", 1)
	graph.AddUndirectedEdge("8", "7", 1)
	graph.AddUndirectedEdge("8", "1", 2)

	result := graph.Dijkstra("1")
	want := map[string]float64{
		"1": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 4,
		"7": 3,
		"8": 2,
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Djikstra result doesnt match!\n Got: %v \n Wanted: %v\n", result, want)
	}

	fmt.Printf("Result: %v", result)

}

func TestUCC(t *testing.T) {
	graph := NewGraph()

	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddVertex("7")
	graph.AddVertex("8")
	graph.AddVertex("9")

	graph.AddUndirectedEdge("1", "2", 1)
	graph.AddUndirectedEdge("1", "3", 1)
	graph.AddUndirectedEdge("1", "4", 1)
	graph.AddUndirectedEdge("5", "6", 1)
	graph.AddUndirectedEdge("6", "7", 1)
	graph.AddUndirectedEdge("8", "9", 1)

	result := graph.UCC()

	component1 := []string{"1", "2", "3", "4"}
	component2 := []string{"5", "6", "7"}
	component3 := []string{"8", "9"}

	for _, node := range component1 {
		if _, ok := result[node]; !ok {
			t.Errorf("Node %s is not present in component 1", node)
		}
	}

	for _, node := range component2 {
		if _, ok := result[node]; !ok {
			t.Errorf("Node %s is not present in component 2", node)
		}
	}

	for _, node := range component3 {
		if _, ok := result[node]; !ok {
			t.Errorf("Node %s is not present in component 3", node)
		}
	}
}

func TestMinHeap(t *testing.T) {

	testHeap := NewMinHeap()
	testHeap.Insert(&Node{distanceFromStart: 100})
	testHeap.Insert(&Node{distanceFromStart: 200})
	testHeap.Insert(&Node{distanceFromStart: 300})
	testHeap.Insert(&Node{distanceFromStart: 400})
	testHeap.Insert(&Node{distanceFromStart: 500})
	testHeap.Insert(&Node{distanceFromStart: 600})

	got := testHeap.ExtractMin().distanceFromStart
	want := 100.0

	if got != want {
		t.Errorf("Node extracted is not smallest! got: %v, want: %v", got, want)

	}

	testNode2 := testHeap.ExtractMin()

	if testNode2.distanceFromStart != 200.0 {
		t.Errorf("Node extracted is not smallest! got: %v, want: %v", testNode2.distanceFromStart, 200.0)
	}

	testNode3 := testHeap.ExtractMin()

	if testNode3.distanceFromStart != 300.0 {
		t.Errorf("Node extracted is not smallest! got: %v, want: %v", testNode3.distanceFromStart, 300.0)
	}
}

// Umgekehrt
func TestMaxHeap(t *testing.T) {

	testHeap := NewMaxHeap()
	testHeap.Insert(&Node{distanceFromStart: 100})
	testHeap.Insert(&Node{distanceFromStart: 200})
	testHeap.Insert(&Node{distanceFromStart: 300})
	testHeap.Insert(&Node{distanceFromStart: 400})
	testHeap.Insert(&Node{distanceFromStart: 500})
	testHeap.Insert(&Node{distanceFromStart: 600})

	got := testHeap.ExtractMax().distanceFromStart
	want := 600.0

	if got != want {
		t.Errorf("Node extracted is not largest! got: %v, want: %v", got, want)

	}

	testNode2 := testHeap.ExtractMax()

	if testNode2.distanceFromStart != 500.0 {
		t.Errorf("Node extracted is not smallest! got: %v, want: %v", testNode2.distanceFromStart, 200.0)
	}

	testNode3 := testHeap.ExtractMax()

	if testNode3.distanceFromStart != 400.0 {
		t.Errorf("Node extracted is not smallest! got: %v, want: %v", testNode3.distanceFromStart, 300.0)
	}
}

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(&Node{ID: "1"})
	queue.Enqueue(&Node{ID: "2"})
	queue.Enqueue(&Node{ID: "3"})
	queue.Enqueue(&Node{ID: "4"})
	queue.Enqueue(&Node{ID: "5"})
	queue.Enqueue(&Node{ID: "6"})

	if queue.IsEmpty() {
		t.Errorf("Queue should already be filled!")
	}

	got := queue.Dequeue().ID
	want := "1"

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}

	got = queue.Dequeue().ID
	want = "2"

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}

}

func TestStack(t *testing.T) {
	stack := NewStack()

	emptyPop := stack.Pop()

	if emptyPop != nil {
		t.Errorf("Stack should be empty")
	}

	stack.Push(&Node{ID: "1"})
	stack.Push(&Node{ID: "2"})
	stack.Push(&Node{ID: "5"})
	stack.Push(&Node{ID: "8"})
	stack.Push(&Node{ID: "3"})

	got := stack.Pop().ID
	want := "3"

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}

	got = stack.Pop().ID
	want = "8"

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}

	got = stack.Pop().ID
	want = "5"

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}

}

func Test_NumVerteces(t *testing.T) {
	graph2 := NewGraph()
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")
	graph2.AddVertex("6")
	graph2.AddVertex("7")
	graph2.AddVertex("8")
	graph2.AddVertex("9")

	graph2.AddDirectedEdge("1", "2", 1)
	graph2.AddDirectedEdge("1", "3", 1)
	graph2.AddDirectedEdge("2", "4", 1)
	graph2.AddDirectedEdge("2", "5", 1)
	graph2.AddDirectedEdge("3", "6", 1)
	graph2.AddDirectedEdge("3", "7", 1)
	graph2.AddDirectedEdge("4", "8", 1)
	graph2.AddDirectedEdge("4", "9", 1)

	got := graph2.NumVertices()
	want := 9

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func Test_NumEdgesDirected(t *testing.T) {
	graph2 := NewGraph()
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")
	graph2.AddVertex("6")
	graph2.AddVertex("7")
	graph2.AddVertex("8")
	graph2.AddVertex("9")

	graph2.AddDirectedEdge("1", "2", 1)
	graph2.AddDirectedEdge("1", "3", 1)
	graph2.AddDirectedEdge("2", "4", 1)
	graph2.AddDirectedEdge("2", "5", 1)
	graph2.AddDirectedEdge("3", "6", 1)
	graph2.AddDirectedEdge("3", "7", 1)
	graph2.AddDirectedEdge("4", "8", 1)
	graph2.AddDirectedEdge("4", "9", 1)

	got := graph2.NumEdges()
	want := 8

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func Test_NumEdgesUndirected(t *testing.T) {
	graph2 := NewGraph()
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")
	graph2.AddVertex("6")
	graph2.AddVertex("7")
	graph2.AddVertex("8")
	graph2.AddVertex("9")

	graph2.AddUndirectedEdge("1", "2", 1)
	graph2.AddUndirectedEdge("1", "3", 1)
	graph2.AddUndirectedEdge("2", "4", 1)
	graph2.AddUndirectedEdge("2", "5", 1)
	graph2.AddUndirectedEdge("3", "6", 1)
	graph2.AddUndirectedEdge("3", "7", 1)
	graph2.AddUndirectedEdge("4", "8", 1)
	graph2.AddUndirectedEdge("4", "9", 1)
	graph2.AddUndirectedEdge("3", "9", 1)
	graph2.AddUndirectedEdge("2", "9", 1)

	got := graph2.NumEdges()
	want := 20

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func TestBFS(t *testing.T) {

	// graph from https://medium.com/basecs/breaking-down-breadth-first-search-cebe696709d9
	graph2 := NewGraph()
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")
	graph2.AddVertex("6")
	graph2.AddVertex("7")
	graph2.AddVertex("8")
	graph2.AddVertex("9")

	graph2.AddDirectedEdge("1", "2", 1)
	graph2.AddDirectedEdge("1", "3", 1)
	graph2.AddDirectedEdge("2", "4", 1)
	graph2.AddDirectedEdge("2", "5", 1)
	graph2.AddDirectedEdge("3", "6", 1)
	graph2.AddDirectedEdge("3", "7", 1)
	graph2.AddDirectedEdge("4", "8", 1)
	graph2.AddDirectedEdge("4", "9", 1)

	result := graph2.BFS("1")

	layer1 := []string{"2", "3"}
	layer2 := []string{"4", "5", "6", "7"}
	layer3 := []string{"8", "9"}

	for _, node := range layer1 {
		if result[node] != 1 {
			t.Errorf("Node %s is not present in layer 1. Instead in layer: %v", node, result[node])
		}
	}

	for _, node := range layer2 {
		if result[node] != 2 {
			t.Errorf("Node %s is not present in layer 2. Instead in layer: %v", node, result[node])
		}
	}

	for _, node := range layer3 {
		if result[node] != 3 {
			t.Errorf("Node %s is not present in layer 3. Instead in layer: %v", node, result[node])
		}
	}
}

func TestDFS(t *testing.T) {

	graph2 := NewGraph()
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")
	graph2.AddVertex("6")
	graph2.AddVertex("7")
	graph2.AddVertex("8")
	graph2.AddVertex("9")
	graph2.AddVertex("10")
	graph2.AddVertex("11")

	graph2.AddDirectedEdge("1", "2", 1)
	graph2.AddDirectedEdge("1", "3", 1)
	graph2.AddDirectedEdge("2", "4", 1)
	graph2.AddDirectedEdge("2", "5", 1)
	graph2.AddDirectedEdge("3", "6", 1)
	graph2.AddDirectedEdge("3", "7", 1)
	graph2.AddDirectedEdge("8", "9", 1)
	graph2.AddDirectedEdge("9", "10", 1)

	result := graph2.DFS("1")

	unreachable := []string{"8", "9", "10"}
	reachable := []string{"1", "2", "3", "4", "5", "6", "7"}

	for _, node := range unreachable {
		if result[node] {
			t.Errorf("Node %s is reachable, should be unreachable.", node)
		}
	}

	for _, node := range reachable {
		if !result[node] {
			t.Errorf("Node %s is unreachable, should be reachable.", node)
		}
	}

}

func TestGoogleWebgraph(t *testing.T) {

	graphGoogle := initWebgraph(t)

	/**
		https://snap.stanford.edu/data/web-Google.html
		Nodes	875713
	    Edges	5105039
		**/

	numEdges := graphGoogle.NumEdges()
	numNodes := graphGoogle.NumVertices()

	if numEdges != 5105039 {
		fmt.Printf("Number of edges should be 5105039 but got %v", numEdges)
	}

	if numNodes != 875713 {
		fmt.Printf("Number of edges should be 875713 but got %v", numNodes)
	}

}

/**
func TestTopo(t *testing.T) {
	// For graph 2: https://www.geeksforgeeks.org/topological-sorting/
	// Expected result for Topo: 5 4 2 3 1 0

	graph2 := NewGraph()
	graph2.AddVertex("0")
	graph2.AddVertex("1")
	graph2.AddVertex("2")
	graph2.AddVertex("3")
	graph2.AddVertex("4")
	graph2.AddVertex("5")

	graph2.AddDirectedEdge("5", "0", 1)
	graph2.AddDirectedEdge("4", "0", 1)
	graph2.AddDirectedEdge("4", "1", 1)
	graph2.AddDirectedEdge("5", "2", 1)
	graph2.AddDirectedEdge("2", "3", 1)
	graph2.AddDirectedEdge("3", "1", 1)

	got := graph2.TopoSort()
	want := map[string]int{
		"0": 6,
		"1": 5,
		"2": 3,
		"3": 4,
		"4": 2,
		"5": 1,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopoSort: Result map is not equal to expected map!")
	}
} **/
