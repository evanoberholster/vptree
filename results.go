package vptree

// Results are the VPTree search results
type Results struct {
	items []uint64
	dist  []uint8
}

func newResults(k int) *Results {
	return &Results{
		items: make([]uint64, 0, k),
		dist:  make([]uint8, 0, k),
	}
}

// Array returns VPTree search results as 2 arrays.
// arrays are always equally lengthed.
func (r Results) Array() ([]uint64, []uint8) {
	return r.items, r.dist
}

// Top returns the largest distance between
func (r Results) Top() (item uint64, dist uint64) { return r.items[0], uint64(r.dist[0]) }

func (r Results) Len() int           { return len(r.items) }
func (r Results) Less(i, j int) bool { return r.dist[i] > r.dist[j] }
func (r Results) Swap(i, j int) {
	r.dist[i], r.dist[j] = r.dist[j], r.dist[i]
	r.items[i], r.items[j] = r.items[j], r.items[i]
}

// Push adds an item and dist to the underlying Results
func (r *Results) Push(item uint64, dist uint64) {
	r.dist = append(r.dist, uint8(dist))
	r.items = append(r.items, item)
	r.up(r.Len() - 1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (r *Results) Pop() (item uint64, dist uint64) {
	n := r.Len() - 1
	r.Swap(0, n)
	r.down(0, n)
	n = r.Len()
	item = r.items[n-1]
	r.items = r.items[0 : n-1]
	dist = uint64(r.dist[n-1])
	r.dist = r.dist[0 : n-1]
	return item, dist
}

func (r *Results) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !r.Less(j, i) {
			break
		}
		r.Swap(i, j)
		j = i
	}
}

func (r *Results) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && r.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !r.Less(j, i) {
			break
		}
		r.Swap(i, j)
		i = j
	}
	return i > i0
}
