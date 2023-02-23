# [Maps](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps)

* don't ask me why they put this in `package main`
* Turn the `dictionary map[string]string` into its own type
* Also turn `Search` into a receiver function
* A common trend I'm seeing right now is to create a variable for a specific error you might expects, like

```go
var ErrNotFound = errors.New("could not find the word you were looking for")
```

* Maps are pointers to `runtime.hmap` structs
* Empty maps are `nil`, which has a gotcha
* Reading from a `nil` map behaves like an empty map but writing to one will cause a runtime panic
* Therefore, never initialize an empty map like `var m map[string]string`
* Instead, do the following

```go
var dict = map[string]string{}
// OR
var dict = make(map[string]string)
```

* Go has a built-in function for deleting entries in hashmaps called `delete(m map[T]K, key T)` 
