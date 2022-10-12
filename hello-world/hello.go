package main

import (
	"fmt"
)

type Language string

const Spanish Language = "Spanish"
const Bahasa Language = "Bahasa"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const bahasaHelloPrefix = "Halo"

func Hello(name string, lang Language) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(lang), name)
}

func greetingPrefix(lang Language) (prefix string) {
	switch lang {
	case Spanish:
		prefix = spanishHelloPrefix
	case Bahasa:
		prefix = bahasaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
