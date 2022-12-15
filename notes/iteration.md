# [Iteration](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration)
* All loops/iterations in go are done with the `for` keyword
* Same shit
  * write test
  * write enough code to make the test run
  * write enough code to make the test pass
```go
package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
```
* We can introduce a constant for the number of iterations so it doesn't re-evaluate each time the func runs
* Also using the `+=` operator now

## Benchmarking
* Writing benchmarks are very similar to writing tests
```go
package iteration

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```
* Very similar to tests indeed
* `b.N` determines how many times the code will run, which is completely left up to the benchmarker
* To run the benchmark, run `go test -bench=.`
* benchmark returns something like 
```
goos: darwin
goarch: amd64
pkg: hello/iteration
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkRepeat-16      11747678               102.7 ns/op
PASS
ok      hello/iteration 1.427s
```
* Key takeaway here is that the code ran 11,747,678 times and the function takes and average of 102.7 nanosecods to run
* benchmarks are run sequentially by default

## Practice
* Change the test so a caller can specify how many times the character is repeated and then fix the code
* Write `ExampleRepeat` to document your function
* Have a look through the strings package. Find functions you think could be useful and experiment with them by writing 
tests like we have here. Investing time learning the standard library will really pay off over time.

## note
* I'm going to be very sparse with my notes going forward since I already know most of the go concepts being taught here
* I'm mostly just using this as an exercise in learning tests