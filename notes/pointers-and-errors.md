# [Pointers & errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors)
* um they're trying to demonstrate why pointer receivers are important but I think Go is doing some optimization??
* Like they're acting as pointer receivers
* Oh lol n/m I was in the wrong directory
* Yeah they're at different memory locations now...
* Yeah now we're chillin
* Making a pointer of a primitive type here
  * downside here is that we have to pass `Bitcoin(10)` to the `Deposit` method now
  * > This can be very useful when you want to add some domain specific functionality on top of existing types.
* We can implement the `Stringer` interface to customize how a type gets printed
* To do this, simply add a `String() string` method to the type

```go
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```
* I'm noticing that we don't have to convert the `b` to `int`...
* To see this change, we need to tell the format string to format the bitcoin as a string (`%s`)

## Errors
* What happens if you withdraw an amount more than balance?
* Return an error!!!!
* Hopefully gonna get some good knowledge on how to write test expecting errors
* Pretty much return a `t.Error` if `err == nil`
  * makes sense
* It's also useful to make sure you get the error _message_ you were expecting
* We used `t.Fatal` on the error nil check because we don't want to try and use the error any further
  * `t.Fatal` will end the error right there
  * calling `err.Error()` on nil will cause a runtime panic!!!
* 
