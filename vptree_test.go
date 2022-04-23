package vptree

import (
	"bufio"
	"encoding/binary"
	"os"
	"testing"
)

const (
	testCount        = 1000000
	testQuery uint64 = 9499926864262602905
	testTau   uint64 = 20
	testK            = 10
)

func readDB() []uint64 {
	f, err := os.Open("pHash.db")
	if err != nil {
		panic(err)
	}
	s, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := s.Size()

	buf := bufio.NewReader(f)

	items := make([]uint64, size/8)
	item := make([]byte, 8)
	for i := 0; i < int(size/8); i++ {
		buf.Read(item)
		items[i] = binary.LittleEndian.Uint64(item)
	}
	return items
}

var vp *VPTree
var arr []uint64

func init() {
	vp = NewVPTree()
	arr = readDB()
	for i := 0; i < len(arr); i++ {
		vp.Add(arr[i])
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	query2 := uint64(16422160614500647000)
	tau := uint64(4)

	b.ResetTimer()
	b.ReportAllocs()
	b.Run("2", func(b *testing.B) {
		tau = 2
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("4", func(b *testing.B) {
		tau = 4
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("6", func(b *testing.B) {
		tau = 6
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("8", func(b *testing.B) {
		tau = 8
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("10", func(b *testing.B) {
		tau = 10
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("12", func(b *testing.B) {
		tau = 12
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
	b.Run("14", func(b *testing.B) {
		tau = 14
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query2, tau, testK)
		}
	})
}

func BenchmarkSearch(b *testing.B) {
	query2 := uint64(16422160614500647000)
	tau := uint64(4)

	b.ResetTimer()
	b.ReportAllocs()
	b.Run("2", func(b *testing.B) {
		tau = 2
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("4", func(b *testing.B) {
		tau = 4
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("6", func(b *testing.B) {
		tau = 6
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("8", func(b *testing.B) {
		tau = 8
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("10", func(b *testing.B) {
		tau = 10
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("12", func(b *testing.B) {
		tau = 12
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
	b.Run("14", func(b *testing.B) {
		tau = 14
		for i := 0; i < b.N; i++ {
			vp.Search(query2, tau, testK)
		}
	})
}

// BenchmarkSearch/16-12 	   27498	     43427 ns/op	    1544 B/op	      71 allocs/op

// BenchmarkSearch/16-12 	   23143	     49293 ns/op	    1400 B/op	      62 allocs/op

// BenchmarkSearch/16-12 	    1249	    986074 ns/op	    1848 B/op	      90 allocs/op

// BenchmarkSearch/16-12 	    1216	    983878 ns/op	     656 B/op	      11 allocs/op

// BenchmarkSearch/16-12 	    1212	    969849 ns/op	     176 B/op	       4 allocs/op

// 2048
// BenchmarkSearch/16-12 	    1262	    930683 ns/op	     160 B/op	       4 allocs/op

// 4096
// BenchmarkSearch/16-12 	    1340	    905357 ns/op	     160 B/op	       4 allocs/op

// 1024
// BenchmarkSearch/16-12 	    1098	    948374 ns/op	     160 B/op	       4 allocs/op

// 512
// BenchmarkSearch/16-12 	    1212	    979466 ns/op	     160 B/op	       4 allocs/op

// 256
// BenchmarkSearch/16-12 	     964	   1037699 ns/op	     160 B/op	       4 allocs/op

// 128
// BenchmarkSearch/16-12 	     994	   1211728 ns/op	     160 B/op	       4 allocs/op

//cpu: AMD Ryzen 5 5600X 6-Core Processor
//BenchmarkSearch/16-12   	    1167	   1045384 ns/op	     160 B/op	       4 allocs/op
//BenchmarkSimStore/16-12 	      79	  14379709 ns/op	    2816 B/op	      94 allocs/op
