package main

type Stack struct {
	nodes []*Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes, n)
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	index := len(s.nodes) - 1
	result := s.nodes[index]
	s.nodes = s.nodes[:index]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}
