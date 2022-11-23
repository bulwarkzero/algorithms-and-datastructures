package union

import graph "new-way/4-1-1-graph"

type Union struct {
	mappings map[*graph.Node]int
	items    []int
	ranks    map[int]int
}

// Union put a and b in same component
func (u *Union) Union(a, b *graph.Node) {
	aComponent := u.Find(a)
	bComponent := u.Find(b)

	if aComponent == bComponent {
		return
	}

	aComponentSize := u.rank(aComponent)
	bComponentSize := u.rank(bComponent)
	largestComponent := aComponent
	smallestComponent := bComponent

	if bComponentSize > aComponentSize {
		largestComponent = bComponent
		smallestComponent = aComponent
	}

	// smallest component merges into largest
	u.items[smallestComponent] = largestComponent
	u.incRank(largestComponent)
}

// Find union component of a node
func (u *Union) Find(node *graph.Node) int {
	rootIndex := u.mappings[node]
	originalItemIndex := rootIndex

	for u.items[rootIndex] != rootIndex {
		rootIndex = u.items[rootIndex]
	}

	// path compression
	for originalItemIndex != rootIndex {
		next := u.items[originalItemIndex]
		u.items[originalItemIndex] = rootIndex
		originalItemIndex = next
	}

	return rootIndex
}

// incRank increment a union component size
func (u *Union) incRank(id int) int {
	if _, ok := u.ranks[id]; ok {
		u.ranks[id]++

		return u.ranks[id]
	}

	u.ranks[id] = 1

	return 1
}

// rank get a union component size
func (u *Union) rank(id int) int {
	if rank, ok := u.ranks[id]; ok {
		return rank
	}

	return 0
}

// New initialize new union
func New(items []*graph.Node) *Union {
	mappings := make(map[*graph.Node]int, len(items))
	unionItems := make([]int, len(items))

	for i, item := range items {
		mappings[item] = i
	}

	for i := range items {
		unionItems[i] = i
	}

	return &Union{mappings, unionItems, make(map[int]int)}
}
