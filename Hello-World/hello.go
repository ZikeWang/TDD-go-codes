package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return choosePrefix(language) + name
}

func choosePrefix(language string) (prefix string){
	switch language{
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
	// this is different from C, we can just call 'return' rather than 'return prefix'
	// to returns the return value of the function we defined (i.e. prefix)
}

func main() {
	fmt.Println(Hello("world", ""))
}
