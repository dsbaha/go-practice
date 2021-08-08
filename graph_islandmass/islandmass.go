package graph_islandmass

import (
//"fmt"
)

func maxAreaOfIsland(grid [][]int) (ret int) {
	g := &Graph{}
	g.BuildGraph(grid)

	cached := make(map[int]bool)

	for _, v := range g.Nodes {
		if _, ok := cached[v.Key]; ok {
			continue
		}

		seen := v.DFS(nil)
		//fmt.Println(seen)
		if len(seen) > ret {
			ret = len(seen)
		}

		for i, j := range seen {
			cached[i] = j
		}

	}

	return
}

func (n *Node) DFS(seen map[int]bool) map[int]bool {
	if seen == nil {
		seen = make(map[int]bool)
	}

	seen[n.Key] = true

	for _, node := range n.Linked() {
		if _, ok := seen[node.Key]; ok {
			//fmt.Println("seen ", node.Key)
			continue
		}

		seen[node.Key] = true
		ret := node.DFS(seen)
		for k, v := range ret {
			seen[k] = v
		}
	}

	return seen
}

func (g *Graph) BuildGraph(grid [][]int) {
	for k, v := range grid {
		for i, j := range v {

			// No Land Mass, Continue On
			if j == 0 {
				g.Water += 1
				continue
			}

			g.Land += 1

			// Key or ID of the Land Mass
			key := k*100 + i

			//Land Mass Found, Process
			node := &Node{
				Key:   key,
				Edges: make(map[int]*Node),
			}

			// Add Node to Grid
			g.AddNode(node)

			//Look North or West, no need for south or east.
			g.AddEdge(node, g.Nodes[key-1])
			g.AddEdge(node, g.Nodes[key-100])
		}
	}
}

func (n *Node) Linked() (ret []*Node) {
	for _, v := range n.Edges {
		ret = append(ret, v)
	}
	return
}

type Graph struct {
	Nodes map[int]*Node
	Land  int
	Water int
}

type Node struct {
	Key   int
	Edges map[int]*Node
}

func (g *Graph) AddNode(n *Node) {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*Node)
	}

	g.Nodes[n.Key] = n
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}

	v1 := g.Nodes[n1.Key]
	v2 := g.Nodes[n2.Key]

	if v1 == nil || v2 == nil {
		return
	}

	// Link Nodes
	v1.Edges[v2.Key] = v2
	v2.Edges[v1.Key] = v1
}
