# vptree

### Description
In-memory VPTree for efficient searching of uint64 hash. Designed to be used with [Perceptual Hashing](https://en.wikipedia.org/wiki/Perceptual_hashing). Useful for searching Image hashes produced by [github.com/corona10/goimagehash](https://github.com/corona10/goimagehash)

## TODO:
[ ] Write tests
[ ] Improve benchmarks

## Benchmarks
 LinearSearch compared with VPTree Search on dataset of ~20000. VPTree search will improve compared to linear Search with decreased tau
 and increased number of items.
```go
    name          old time/op  new time/op  delta
    Search/2-12   22.3µs ± 3%   7.1µs ± 6%  -68.30%  (p=0.000 n=20+20)
    Search/4-12   22.4µs ± 6%  10.7µs ± 6%  -52.11%  (p=0.000 n=20+19)
    Search/6-12   22.6µs ± 4%  15.9µs ± 5%  -29.71%  (p=0.000 n=20+20)
    Search/8-12   22.9µs ± 5%  18.0µs ± 5%  -21.40%  (p=0.000 n=20+19)
    Search/10-12  22.8µs ± 5%  19.6µs ± 5%  -13.79%  (p=0.000 n=20+20)
    Search/12-12  22.6µs ± 4%  19.7µs ± 4%  -12.88%  (p=0.000 n=20+20)
    Search/14-12  22.8µs ± 4%  20.0µs ± 5%  -12.14%  (p=0.000 n=19+20)
```

Search Benchmarks compared for 1,000,000 items:
```go
//cpu: AMD Ryzen 5 5600X 6-Core Processor
//BenchmarkSearch/16-12   	    1167	   1045384 ns/op	     160 B/op	       4 allocs/op
//BenchmarkSimStore/16-12 	      79	  14379709 ns/op	    2816 B/op	      94 allocs/op
```

## Inspired by
[github.com/dgryski/go-simstore](https://github.com/dgryski/go-simstore)

## Contributing
Please create a PR. Would appreciate contributions.

## Authors
Evan Oberholster (2022)

## License
MIT License