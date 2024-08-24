package main

import "fmt"

const englishHelloPrefix = "Hello, "
const turkishHelloPrefix = "Merhaba, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Turkish":
		return turkishHelloPrefix + name
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Ilhan", turkishHelloPrefix))
}
