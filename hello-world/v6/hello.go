package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "你好， "

// Hello returns a personalised greeting in a given language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	if language == chinese {
		return chineseHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
	fmt.Println(Hello("世界", chinese))
}
