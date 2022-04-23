package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/evanoberholster/vptree"
)

func main() {
	count := 10000000

	vp := vptree.NewVPTree()
	arr := make([]uint64, 0, count)

	for i := 0; i < count; i++ {
		item := rand.Uint64()
		arr = append(arr, item)
		vp.Add(item)
	}
	fmt.Println(vp)
	query := uint64(9716099646001308040)
	tau := uint64(3)
	k := 10

	start := time.Now()
	res, distances := vp.Search(query, tau, k).Array()
	elapsed := time.Since(start)
	start = time.Now()
	res2, distances2 := vptree.LinearSearch(arr, query, tau, k).Array()
	elapsed2 := time.Since(start)

	fmt.Println("VPTree Results:", len(res), elapsed)
	for i, r := range res {
		fmt.Println("r: ", r, "  \tdist:", distances[i])
	}
	fmt.Println("\nLinear Results:", len(res2), elapsed2)
	for i, r := range res2 {
		fmt.Println("r: ", r, "  \tdist:", distances2[i])
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
