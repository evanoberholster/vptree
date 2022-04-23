// Package vptree is an implementation of Vantage Point Trees in Go for
// bit hash of uint64.
// It is ideally suited for image similarity using an uint64 hash
package vptree

import "sort"

// VPTree is a vantage point tree
type VPTree struct {
	root *node
}

const (
	// itemsLimit is the limit of items per node before a split is called
	itemsLimit = 512
)

// NewVPTree returns a new VPTree
func NewVPTree() *VPTree {
	return &VPTree{
		root: &node{},
	}
}

// Add adds an item to the given VPTree
func (vp *VPTree) Add(item uint64) {
	vp.root.add(item)
}

// Search a VPTree given the following query, tau, and k
func (vp *VPTree) Search(query uint64, tau uint64, k int) (r *Results) {
	r = newResults(k)

	vp.root.search(query, &tau, k, r)

	sort.Sort(sort.Reverse(r))

	return
}

// Search an arr using Linear Search with the given query, tau, and k
func LinearSearch(arr []uint64, query uint64, tau uint64, k int) (r *Results) {
	r = newResults(k)

	for _, item := range arr {
		d := distance(item, query)
		if d <= tau {
			if r.Len() == k {
				r.Pop()
			}
			r.Push(item, d)
			if r.Len() == k {
				_, tau = r.Top()
			}
		}
	}

	sort.Sort(sort.Reverse(r))

	return
}
