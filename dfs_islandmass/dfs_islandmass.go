package dfs_islandmass

func maxAreaOfIsland(grid [][]int) (ret int) {
	dfs := &DFS{
		cache: make(map[int]bool),
		grid:  grid,
	}

	for i := range dfs.grid {
		var val int
		for j := range dfs.grid[i] {
			val = dfs.area(i, j)
			
			if val > ret {
				ret = val
			}
		}
	}

	return
}

type DFS struct {
	cache map[int]bool
	grid  [][]int
}

// depth first search algo
func (d *DFS) area(a, b int) (ret int) {
	id := a*100 + b

	if _, ok := d.cache[id]; ok {
		return
	}

	if a < 0 || a > len(d.grid)-1 {
		return
	}

	if b < 0 || b > len(d.grid[0])-1 {
		return
	}

	if d.grid[a][b] == 0 {
		return
	}

	d.cache[id] = true

	ret += 1
	ret += d.area(a+1, b)
	ret += d.area(a-1, b)
	ret += d.area(a, b+1)
	ret += d.area(a, b-1)
	return
}
