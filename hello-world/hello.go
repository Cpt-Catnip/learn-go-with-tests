package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(n string, l string) string {
	if n == "" {
		n = "World"
	}

	p := greetingPrefix(l)

	return p + n
}

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

func main() {
	fmt.Println(Hello("world", ""))
}
