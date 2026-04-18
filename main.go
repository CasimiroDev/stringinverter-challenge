package main

import (
	"fmt"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i++ {
		runes[i], runes[j] = runes[j], runes[i]
		j--
	}

	return string(runes)
}

func main() {
	fmt.Println(ReverseString("golang"))
	fmt.Println(ReverseString("olá"))
	fmt.Println(ReverseString("mundo! 🇧🇷"))
}
