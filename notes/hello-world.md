# [Hello, world](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)

* you need to initialize a module in order to run test
* this is done using `go mod init <MODULENAME>`
  * in our case, `go mod init hello`
* here's what our simple module looks like

```go
// hello.go
package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

* Now we can simply run `go test` to run our test file

```bash
$ go test
PASS
ok  	hello	0.508s
```

* this is what a failing test looks like

```bash
$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:10: got "Hello, world" want "Hello, worl"
FAIL
exit status 1
FAIL	hello	0.507s
```

* writing tests follows the same rules as writing Go code with a few additions
  * files need to have the format `*_test.go`
  * test functions need to start with the word `Test`
  * the test function only takes one argument `t *testing.T`
  * need to `import "testing"` to using `*testing.T`

* we added the ability to pass in a name argument, but we can refactor the return statement to use a `constant`!
	* woo I haven't used these yet!

```go
const englishHelloPrefix = "Hello, "

func Hello(n string) string {
	return englishHelloPrefix + n
}
```

* using constants improves performance
	* very negligible in this example, but you get the idea
* Now introducing default arguments
* First the test, though

```go
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
```

* Look, subtests!
* Note that we also refactored re-used logic in the test; refactoring isn't just for prod code
* `testing.TB` is an interface that both `*testing.T` and `*testing.B` satisfy
	* passing this lets us use helper functions within a test or benchmark
* `t.Helper()` tells the test suite that the function is a helper function and not a test itself
* this way, if a test fails inside that function, we get the line number of the failure in the test and not the helper
* Okay lets add a language parameter

```go
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(n string, l string) string {
	if n == "" {
		n = "World"
	}

	if l == "Spanish" {
		return spanishHelloPrefix + n
	}
	if l == "French" {
		return frenchHelloPrefix + n
	}
	return englishHelloPrefix + n
}
```

* now that we have a lot of "cases" for the language parameter `l`, we can leverage the `switch` control flow statement

```go
func greetingPrefix(l string) (p string) {
	switch l {
	case french:
		p = frenchHelloPrefix
	case spanish:
		p = spanishHelloPrefix
	default:
		p = englishHelloPrefix
	}

	return
}
```

* Some more new stuff! Love it.
* named return value: `p string`
	* named return values will be initialized to the zero value of the specified type (`""` in this case)
	* this also _creates_ a variable for us to use in the function
		* this probably means we _always_ have to consider default values for named return values
	* you only have to use the `return` keyword without specifying what you're returning since you already did that in the function signature
	* These will display in the go doc, which makes things clearer
