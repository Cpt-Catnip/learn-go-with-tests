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
