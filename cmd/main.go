package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/evanoberholster/vptree"
)

func main() {
	count := 1000000

	vp := vptree.NewVPTree()
	arr := make([]uint64, 0, count)

	for i := 0; i < count; i++ {
		item := rand.Uint64()
		arr = append(arr, item)
		vp.Add(item)
	}

	query := uint64(15480771490665767304)
	tau := uint64(0)
	k := 10

	// Run VP Search
	start := time.Now()
	res, distances := vp.Search(query, tau, k).Array()
	elapsed := time.Since(start)

	fmt.Println("VPTree Results:", len(res), elapsed)
	for i, r := range res {
		fmt.Println("r: ", r, "  \tdist:", distances[i])
	}

	// Run Linear Search
	start = time.Now()
	res, distances = vptree.LinearSearch(arr, query, tau, k).Array()
	elapsed = time.Since(start)

	fmt.Println("\nLinear Results:", len(res), elapsed)
	for i, r := range res {
		fmt.Println("r: ", r, "  \tdist:", distances[i])
	}

	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
