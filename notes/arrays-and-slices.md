# [Arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
```go
package main

// sum_test.go
func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// sum.go
func Sum(numbers [5]int) int {
	return 0
}
```
* note that we're making the parameter type `[5]int`
* That's probably not the best way to do this
* update the tests so the function expects an arbitrary-length array
* Hmm... one of the tests is being passed `[5]int` and the other is being passed `[]int`
  * let's see how this plays out
* Oh, the solution is to change the first test
  * There's no way to make a function that accepts a slice _and_ an array
* It's not about having as many test as you can, it's about having as much __confidence__ as you can
* If you can cover all you cases with fewer tests, do it.
* Here, our two tests are redundant
* You can run `go test -cover` to get a coverage report of your code
* Output will look something like
```
PASS
coverage: 100.0% of statements
ok      hello/arrays    0.109s
```
* Note that this looks fine even though we have redundant tests >.>
* Deleting the other test yields __no change__ in our coverage report

## Sum All
* We want a function `SumAll` that accepts a _variable number of slices_ and returns a slice containing the sum of each passed slice
* Gonna write a variadic function B)
* This is interesting! Slices _can't_ use the equality operator!!!
* It's possible to write a function that iterates over each elem in `got` and `want` but go provides us with [`reflect.DeepEqual`](https://pkg.go.dev/reflect#DeepEqual)
  * `func DeepEqual(x, y any) bool`
* `DeepEqual` is _not_ type safe!
* The code will compile if we pass values with diff types to `x` and `y`

## Sum All Tails
* Now we want a function that does the same thing as `SumAll` except that it ignores the first element of each slice
* Slice syntax!
* `slice[low:high]`
* What happens with empty slices??
* Write a test about it!
* Yep we get an out-of-bounds error. That makes sense.
* Just check the length of `numbers` :^|
* Yay now we're refactoring our test to reuse that check function!
```go
checkSums := func(t testing.TB, got, want []int) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```
* Now we can use that anywhere in the test we want to do that check!
  * I should look up the docs on `t.Helper()` and `testing.TB`
* `checkSums` is also type-safe, so we don't have to worry about `DeepEqual` anymore
