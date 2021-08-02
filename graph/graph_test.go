package graph

import (
	"testing"
)

func TestGraphCreate(t *testing.T) {
	g := NewGraph()
	if g.Nodes == nil {
		t.Error("graph nodes map uninitialized.")
	}
}

func TestNodeCreate(t *testing.T) {
	ID := 1
	n := NewNode(ID)
	if n.Key != ID {
		t.Errorf("node key invalid, expecting %v got %v", ID, n.Key)
	}

	if n.Edges == nil {
		t.Error("node edges map uninitialized.")
	}

	if n.Data == nil {
		t.Error("node data map uninitialized")
	}
}

func CreatePopulatedGraph() (g *Graph, n []*Node) {
	g = NewGraph()
	for i := 0; i < 6; i++ {
		nn := NewNode(i)
		n = append(n, nn)
		g.AddNode(nn)
	}

	for i := 0; i < len(n)-1; i++ {
		g.AddEdge(n[i], n[i+1])
	}

	for i := 0; i < len(n); i++ {
		g.AddEdge(n[0], n[i])
	}

	return
}


func TestLinkedNodeTraversal(t *testing.T) {
	_, nodes := CreatePopulatedGraph()

	for i := 0; i < len(nodes); i++ {
		if len(nodes[i].Linked()) < 2 {
			t.Error("invalid amount of nodes linked.")
		}
	}
}

func TestGraphWithNodes(t *testing.T) {
	g, _ := CreatePopulatedGraph()
	ret := g.String()
	if ret == "" {
		t.Error("invalid graph print")
	}
}

func TestGraphWithDFS(t *testing.T) {
	g, _ := CreatePopulatedGraph()

	if len(g.DFS(g.Nodes[0])) != 5 {
		t.Error("invalid number of nodes")
	}

	if len(g.DFS(g.Nodes[4])) != 3 {
		t.Error("invalid number of nodes")
	}
}