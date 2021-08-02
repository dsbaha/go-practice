// basic practice graph data store in go.
// generally, graphs have nodes and edges.
// add edge (links) and add nodes are generally needed.

package graph

import (
	"fmt"
	"sync"
)

// The Node Struct
// Key, Edges, and Data
// Make it Thread Safe with sync
type Node struct {
	Key   int
	Edges map[int]*Node
	Data  map[string]interface{}
	mtx   sync.RWMutex
}

// The Graph
// Make it Thread Safe with sync
type Graph struct {
	Nodes map[int]*Node
	mtx   sync.RWMutex
	wg    sync.WaitGroup
}

// NewNode initializes a node with specified key.
func NewNode(k int) *Node {
	return &Node{
		Key:   k,
		Edges: make(map[int]*Node),
		Data:  make(map[string]interface{}),
	}
}

// NewGraph creates a new graph!
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

// Gets a Node in a Thread Safe Way
func (g *Graph) GetNode(n *Node) *Node {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	g.mtx.RLock()
	defer g.mtx.RUnlock()

	return g.Nodes[n.Key]
}

// Link Node in a Thread Safe Way
func (n *Node) LinkNode(n1 *Node) {

	n1.mtx.RLock()
	defer n1.mtx.RUnlock()

	if n.Key == n1.Key {
		// Don't Link Itself!
		return
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()

	n.Edges[n1.Key] = n1
}

// Unlink Node in a Thread Safe Way
func (n *Node) UnlinkNode(n1 *Node) {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	n1.mtx.Lock()
	defer n.mtx.Unlock()

	delete(n.Edges, n1.Key)
}

// Returns the Linked Nodes
func (n *Node) Linked() (ret []*Node) {
	for _, v := range n.Edges {
		ret = append(ret, v)
	}
	return
}

// Add a Node to the Graph in a Thread Safe Way
func (g *Graph) AddNode(n *Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	g.Nodes[n.Key] = n
}

// Add Edge in a thread safe way.
func (g *Graph) AddEdge(n1, n2 *Node) {
	v1 := g.GetNode(n1)
	v2 := g.GetNode(n2)

	if v1 == nil || v2 == nil {
		return
	}

	// Link the nodes.
	v1.LinkNode(v2)
	v2.LinkNode(v1)

	// Add to Graph
	g.AddNode(v1)
	g.AddNode(v2)
}

// Remove the Edge in a Thread Safe Way
func (g *Graph) RemoveEdge(n1, n2 *Node) {
	v1 := g.GetNode(n1)
	v2 := g.GetNode(n2)

	if v1 == nil || v2 == nil {
		return
	}

	//Unlink the Nodes
	v1.UnlinkNode(v2)
	v2.UnlinkNode(v1)

	//Remove from Graph
	g.RemoveNode(v1)
	g.RemoveNode(v2)
}

// Remove a node in a thread safe way.
func (g *Graph) RemoveNode(n *Node) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	delete(g.Nodes, n.Key)
}

// Return a string representation of the graph.
func (g *Graph) String() (ret string) {
	ret += "graph links\n"
	for _, v := range g.Nodes {
		ret += fmt.Sprintf("%v -> ", v.Key)
		for _, i := range v.Linked() {
			ret += fmt.Sprintf(" %v", i.Key)
		}
		ret += "\n"
	}
	return
}

// Depth First Search DFS
// Good for searching linked nodes at START point.
func (g *Graph) DFS(start *Node) (ret []*Node) {
	for _, v := range start.Edges {
		ret = append(ret, v)
	}
	return
}
