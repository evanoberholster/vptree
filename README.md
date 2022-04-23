# vptree

### Description
In-memory VPTree for storting and efficient searching of uint64 hash.

## Benchmarks
 LinearSearch compared with VPTree Search on dataset of ~ 20000. VPTree search will improve compared to linear Search with decreased tau
 and increased number of items.
``
    name          old time/op  new time/op  delta
    Search/2-12   22.3µs ± 3%   7.1µs ± 6%  -68.30%  (p=0.000 n=20+20)
    Search/4-12   22.4µs ± 6%  10.7µs ± 6%  -52.11%  (p=0.000 n=20+19)
    Search/6-12   22.6µs ± 4%  15.9µs ± 5%  -29.71%  (p=0.000 n=20+20)
    Search/8-12   22.9µs ± 5%  18.0µs ± 5%  -21.40%  (p=0.000 n=20+19)
    Search/10-12  22.8µs ± 5%  19.6µs ± 5%  -13.79%  (p=0.000 n=20+20)
    Search/12-12  22.6µs ± 4%  19.7µs ± 4%  -12.88%  (p=0.000 n=20+20)
    Search/14-12  22.8µs ± 4%  20.0µs ± 5%  -12.14%  (p=0.000 n=19+20)
``

## Authors
Evan Oberholster