package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 🔁 Переворачивает буквы в каждом слове
func reverseWords(input string) string {
	words := strings.Fields(input)
	var reversedWords []string

	for _, word := range words {
		runes := []rune(word)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		reversedWords = append(reversedWords, string(runes))
	}
	return strings.Join(reversedWords, " ")
}

// 🔄 Меняет местами чётные и нечётные буквы
func swapEvenOddRunes(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i += 2 {
		runes[i], runes[i+1] = runes[i+1], runes[i]
	}
	return string(runes)
}

// 🔒 Шифрует текст: сначала переворот, потом замена букв
func encrypt(text string) string {
	reversed := reverseWords(text)
	words := strings.Fields(reversed)
	for i := range words {
		words[i] = swapEvenOddRunes(words[i])
	}
	return strings.Join(words, " ")
}

// 🔓 Расшифровывает текст: сначала обратно буквы, потом переворот
func decrypt(text string) string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = swapEvenOddRunes(words[i]) // буквы обратно местами
	}
	result := strings.Join(words, " ")
	return reverseWords(result) // теперь переворачиваем
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введи текст: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Выбери действие:")
	fmt.Println("1 - Зашифровать текст")
	fmt.Println("2 - Расшифровать текст")
	fmt.Print("Твой выбор: ")

	choiceInput, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(choiceInput)

	switch choice {
	case "1":
		result := encrypt(input)
		fmt.Println("🔒 Зашифрованный текст:", result)
	case "2":
		result := decrypt(input)
		fmt.Println("🔓 Расшифрованный текст:", result)
	default:
		fmt.Println("❌ Неверный выбор. Попробуй снова.")
	}
}
