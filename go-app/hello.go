package main

import (
	"fmt"
)

const english = "English"
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = getDefaultName(language)
	}

	return getGreetingPrefix(language) + name

}

func getDefaultName(language string) (name string) {
	switch language {
	case english:
		name = "World"
	case spanish:
		name = "Elodie"
	case french:
		name = "monde"
	default:
		name = "World"
	}

	return
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
