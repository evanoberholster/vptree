package vptree

import (
	"bufio"
	"encoding/binary"
	"math/rand"
	"os"
	"testing"
)

const (
	testCount        = 1000000
	query     uint64 = 9716099646001308040
	k                = 10
)

var (
	tau uint64 = 0
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
	//arr = readDB()
	//for i := 0; i < len(arr); i++ {
	//	vp.Add(arr[i])
	//}
	arr = make([]uint64, 0, testCount)
	for i := 0; i < testCount; i++ {
		n := rand.Uint64()
		arr = append(arr, n)
		vp.Add(n)
	}
}

func BenchmarkLinearSearch(b *testing.B) {

	b.ResetTimer()
	b.ReportAllocs()
	b.Run("2", func(b *testing.B) {
		tau = 2
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("4", func(b *testing.B) {
		tau = 4
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("6", func(b *testing.B) {
		tau = 6
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("8", func(b *testing.B) {
		tau = 8
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("10", func(b *testing.B) {
		tau = 10
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("12", func(b *testing.B) {
		tau = 12
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
	b.Run("14", func(b *testing.B) {
		tau = 14
		for i := 0; i < b.N; i++ {
			LinearSearch(arr, query, tau, k)
		}
	})
}

func BenchmarkVPSearch(b *testing.B) {

	b.ResetTimer()
	b.ReportAllocs()
	b.Run("2", func(b *testing.B) {
		tau = 2
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("4", func(b *testing.B) {
		tau = 4
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("6", func(b *testing.B) {
		tau = 6
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("8", func(b *testing.B) {
		tau = 8
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("10", func(b *testing.B) {
		tau = 10
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("12", func(b *testing.B) {
		tau = 12
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
	b.Run("14", func(b *testing.B) {
		tau = 14
		for i := 0; i < b.N; i++ {
			vp.Search(query, tau, k)
		}
	})
}
