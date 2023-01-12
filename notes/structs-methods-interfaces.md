#[Structs, methods & interfaces](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces)

* sort of going over this testing pattern
  * write test
  * write enough code to pass
  * repeat
* Now there's a `Rectangle` type and the functions accept one as a parameter
* What about circles?
* We're using `%g` in our format string now, which just prints a more precise decimal
* Now we need to write our functions to accept a rect or a circ
  * some languages let you do function overloads, but not golang
* To solve this, we're gonna use methods 😎
* To make a helper function that accepts both circles and rectangles, we'll define an interface for `Shape`s
* Creating this interface decouples the helper from the concrete type
* All you need is something that has the method `Area()`

## Table Driven Tests
* This is what I've been waiting for!

```go
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
```

* using a slice of anonymous structs defining what we want from each test and other stuff we need
  * e.g. we need a `Shape` for each area test
* demonstrate benefit by adding a test for triangles

## Some improved error logging
* We can change the error message to `"%#v got %g want %g"` where `%#v` prints out the struct that generated the error
* We can also wrap each test in the loop in a `t.Run` so the errors are even more descriptive
