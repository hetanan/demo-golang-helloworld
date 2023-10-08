package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("hn", "French"))
}
