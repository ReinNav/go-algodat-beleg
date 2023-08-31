package main

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Dequeue() *Node {
	if len(q.nodes) > 0 {
		node := q.nodes[0]
		q.nodes = q.nodes[1:]
		return node
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}
