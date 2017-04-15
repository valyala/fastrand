[![Build Status](https://travis-ci.org/valyala/fastrand.svg)](https://travis-ci.org/valyala/fastrand)
[![GoDoc](https://godoc.org/github.com/valyala/fastrand?status.svg)](http://godoc.org/github.com/valyala/fastrand)
[![Go Report](https://goreportcard.com/badge/github.com/valyala/fastrand)](https://goreportcard.com/report/github.com/valyala/fastrand)


# fastrand

Fast pseudorandom number generator.


# Features

- Optimized for speed.
- Performance scales on multiple CPUs.

# How does it work?

It abuses [sync.Pool](https://golang.org/pkg/sync/#Pool) for maintaining
"per-CPU" pseudorandom number generators.

TODO: firgure out how to use real per-CPU pseudorandom number generators.


# Benchmark results


```
GOMAXPROCS=1 go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/valyala/fastrand
BenchmarkUint32n/1         	50000000	        30.5 ns/op
BenchmarkUint32n/2         	50000000	        30.5 ns/op
BenchmarkUint32n/4         	50000000	        30.6 ns/op
BenchmarkUint32n/8         	50000000	        30.5 ns/op
BenchmarkRNGUint32n        	200000000	         6.51 ns/op
BenchmarkMathRandInt31n/1  	50000000	        31.8 ns/op
BenchmarkMathRandInt31n/2  	50000000	        36.3 ns/op
BenchmarkMathRandInt31n/4  	50000000	        39.6 ns/op
BenchmarkMathRandInt31n/8  	50000000	        41.5 ns/op
BenchmarkMathRandRNGInt31n 	100000000	        18.3 ns/op


GOMAXPROCS=2 go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/valyala/fastrand
BenchmarkUint32n/1-2           	50000000	        38.4 ns/op
BenchmarkUint32n/2-2           	50000000	        39.9 ns/op
BenchmarkUint32n/4-2           	50000000	        37.2 ns/op
BenchmarkUint32n/8-2           	50000000	        37.9 ns/op
BenchmarkRNGUint32n-2          	500000000	         3.37 ns/op
BenchmarkMathRandInt31n/1-2    	30000000	        47.2 ns/op
BenchmarkMathRandInt31n/2-2    	20000000	       160 ns/op
BenchmarkMathRandInt31n/4-2    	10000000	       188 ns/op
BenchmarkMathRandInt31n/8-2    	10000000	       135 ns/op
BenchmarkMathRandRNGInt31n-2   	200000000	         9.31 ns/op


GOMAXPROCS=4 go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/valyala/fastrand
BenchmarkUint32n/1-4           	100000000	        22.8 ns/op
BenchmarkUint32n/2-4           	100000000	        23.4 ns/op
BenchmarkUint32n/4-4           	100000000	        23.1 ns/op
BenchmarkUint32n/8-4           	100000000	        23.2 ns/op
BenchmarkRNGUint32n-4          	500000000	         3.24 ns/op
BenchmarkMathRandInt31n/1-4    	10000000	       136 ns/op
BenchmarkMathRandInt31n/2-4    	10000000	       166 ns/op
BenchmarkMathRandInt31n/4-4    	10000000	       184 ns/op
BenchmarkMathRandInt31n/8-4    	10000000	       194 ns/op
BenchmarkMathRandRNGInt31n-4   	200000000	         8.32 ns/op


GOMAXPROCS=8 go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/valyala/fastrand
BenchmarkUint32n/1-8           	100000000	        17.8 ns/op
BenchmarkUint32n/2-8           	100000000	        18.5 ns/op
BenchmarkUint32n/4-8           	100000000	        18.1 ns/op
BenchmarkUint32n/8-8           	100000000	        18.3 ns/op
BenchmarkRNGUint32n-8          	500000000	         3.28 ns/op
BenchmarkMathRandInt31n/1-8    	10000000	       157 ns/op
BenchmarkMathRandInt31n/2-8    	10000000	       166 ns/op
BenchmarkMathRandInt31n/4-8    	10000000	       185 ns/op
BenchmarkMathRandInt31n/8-8    	10000000	       191 ns/op
BenchmarkMathRandRNGInt31n-8   	200000000	         9.00 ns/op
```
