# [Integers](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers)

* To start things off, we're going to write a _test_ for a function that adds two integers
* Of course, the test fails
* > Write the minimal amount of code for the test to run and check the failing test output
* In this case, that would be something like

```go
package integers

func Add(x, y int) int {
	return 0
}
```

* Trivial and obviously failing, but it lets our test run
  * Here we're using the shorthand that lets us declare two parameters with the same type
* Now the test is failing but at least it can compile so we actually get a helpful error message
* > Write enough code to make it pass
* make it return the sum `x + y`!

## Refactor
* Not really a lot to improve on here, but we can add some documentation!
* Adding comments to functions makes them appear in the go doc.

## Examples
* You can add examples???
* They get compiled and typically live in the `_test.go` file for the package
```go
package integers

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
* to see the output of the example we need to run `go test -v` and the test _needs_ the `Output:` comment
* This code will apparently now appear in the `godoc`
  * neat
* Oh, neat! It really does!
