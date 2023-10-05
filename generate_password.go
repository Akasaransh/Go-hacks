package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_+=<>?{}[]|"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	passwordLength := 12 // Change this to your desired password length
	password := generatePassword(passwordLength)
	fmt.Println("Generated Password:", password)
}

func generatePassword(length int) string {
	charset := uppercaseLetters + lowercaseLetters + digits + specialChars
	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	shuffle(password)
	return string(password)
}

func shuffle(s []byte) {
	n := len(s)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
