package vptree

import (
	"sort"
)

type node struct {
	threshold uint64
	items     []uint64
	left      *node
	right     *node
}

func (n *node) search(q uint64, tau *uint64, k int, r *Results) {
	if n.left == nil {
		for i := 0; i < len(n.items); i++ {
			d := distance(n.items[i], q)
			if d <= *tau {
				if r.Len() == k {
					r.Pop()
				}
				r.Push(n.items[i], d)
				if r.Len() == k {
					_, *tau = r.Top()
				}
			}
		}
	} else {
		d := distance(n.items[0], q)
		if d < n.threshold {
			if d < n.threshold+*tau {
				n.left.search(q, tau, k, r)
			}
			if d >= n.threshold-*tau {
				n.right.search(q, tau, k, r)
			}

		} else {
			if d >= n.threshold-*tau {
				n.right.search(q, tau, k, r)
			}
			if d <= n.threshold+*tau {
				n.left.search(q, tau, k, r)
			}
		}
	}
}

func (n *node) add(item uint64) {
	if n.left == nil {
		if len(n.items) < itemsLimit {
			// deduplicate
			for i := 0; i < len(n.items); i++ {
				if item == n.items[0] {
					return
				}
			}
			// add item
			n.items = append(n.items, item)
			return
		}
		// split node when full
		n.split()
	}
	d := distance(n.items[0], item)
	if d == 0 {
		return // duplicate
	}
	if d < n.threshold {
		n.left.add(item)
	} else {
		n.right.add(item)
	}
}

func (n *node) split() {
	count := len(n.items)
	median := n.items[0]
	distances := make([]uint64, count)
	for i := 1; i < count; i++ {
		distances[i] = distance(median, n.items[i])
	}
	sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })
	n.threshold = distances[count/2]
	n.left = &node{}
	n.right = &node{}

	for i := 1; i < count; i++ {
		if distance(n.items[i], n.items[0]) < n.threshold {
			n.left.add(n.items[i])
		} else {
			n.right.add(n.items[i])
		}
	}
	n.items = n.items[0:1:1]
	//fmt.Println("Split lNode:", len(n.lNode.items), "rNode:", len(n.rNode.items))
}
