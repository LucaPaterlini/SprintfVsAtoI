# Performance Comparison int to string conversion

## Intro

I was just reviewing a solution on [exercism](https://exercism.io) when I found this interesting fact,
my test over a function containing a fmt.Sprintf was slower than strconv.AtoI,
while the student was having the other way around by more than the double margin.

## Wish
So I have written this package that just to the comparison,
would you mind sharing in the issues what your results are and motivate why they are this way?

## My results
```
goos: linux
goarch: amd64
BenchmarkIntegerToASCII-4    	10000000	       120 ns/op	      16 B/op	       2 allocs/op
BenchmarkIntegerToAscii2-4   	50000000	        39.5 ns/op	       7 B/op	       0 allocs/op
PASS
```

My CPU is a i5 2 real cores 4 threaded.

