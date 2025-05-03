package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Character sets for password generation
const lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
const uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specialChars = "!@#$%^&*-{};'.<>/?`"

func main() {
	// Flags to determine character types to include
	var useLetters, useNumbers, useSpecialChars, useUppercaseLetters bool

	// Temporary string inputs from user
	var useLetters_, useNumbers_, useSpecialChars_, useUppercaseLetters_ string

	// Ask for password length
	fmt.Println("Enter the size of your password")
	var length int
	fmt.Scanln(&length)

	// Ask if lowercase letters should be included
	fmt.Println("do you want lowerCase in your password : y/n")
	fmt.Scanln(&useLetters_)
	if useLetters_ == "y" || useLetters_ == "Y" {
		useLetters = true
	}

	// Ask if numbers should be included
	fmt.Println("do you want numbers in your password : y/n")
	fmt.Scanln(&useNumbers_)
	if useNumbers_ == "y" || useNumbers_ == "Y" {
		useNumbers = true
	}

	// Ask if special characters should be included
	fmt.Println("do you want special characters in your password : y/n")
	fmt.Scanln(&useSpecialChars_)
	if useSpecialChars_ == "y" || useSpecialChars_ == "Y" {
		useSpecialChars = true
	}

	// Ask if uppercase letters should be included
	fmt.Println("do you want upperCase letters in your password: y/n")
	fmt.Scanln(&useUppercaseLetters_)
	if useUppercaseLetters_ == "y" || useUppercaseLetters_ == "Y" {
		useUppercaseLetters = true
	}

	// Generate and print the password
	password := passwordGenerate(length, useLetters, useUppercaseLetters, useNumbers, useSpecialChars)
	fmt.Println("Generated password:", password)
}

// Generates password of specified length using selected character sets
func passwordGenerate(length int, useLetter bool, uppercaseletters bool, useNumbers bool, useSpecial bool) string {
	var characterSet string

	// Build the character set based on user input
	if useLetter {
		characterSet += lowercaseLetters
		if uppercaseletters {
			characterSet += uppercaseLetters
		}
	}
	if useNumbers {
		characterSet += numbers
	}
	if useSpecial {
		characterSet += specialChars
	}

	// If no character set was selected, terminate
	if len(characterSet) == 0 {
		panic("Please select at least one character set")
	}

	// Create a byte slice for the password
	bytes := make([]byte, length)

	// Generate initial password
	bt := tryingPasswords(bytes, characterSet)

	// For very short passwords, regenerate anyway
	if length <= 2 {
		tryingPasswords(bytes, characterSet)
	}

	// These conditions attempt to enforce inclusion of each selected type
	// However, due to logic bugs (e.g., missing assignment to bt), it may not work as intended

	if useLetter && useNumbers && useSpecial && uppercaseletters {
		for checkLetters(bt) == false || checkNumbers(bt) == false || checkSpecials(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet) // fixed: update bt
		}
	}
	if useLetter && useNumbers && uppercaseletters {
		for checkLetters(bt) == false || checkNumbers(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet)
		}
	}
	if useLetter && useSpecial && uppercaseletters {
		for checkLetters(bt) == false || checkSpecials(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet)
		}
	}
	if useNumbers && useSpecial && uppercaseletters {
		for checkNumbers(bt) == false || checkSpecials(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet)
		}
	}
	if useSpecial && uppercaseletters {
		for checkSpecials(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet)
		}
	}
	if useNumbers && uppercaseletters {
		for checkNumbers(bt) == false || checkUpperCase(bt) == false {
			bt = tryingPasswords(bytes, characterSet)
		}
	}

	// Return the generated password as a string
	return string(bytes)
}

// Fills the byte slice with random characters from the character set
func tryingPasswords(bytes []byte, characterSet string) []byte {
	source := rand.NewSource(time.Now().UnixNano()) // seed RNG with current time
	r := rand.New(source)

	for i := range bytes {
		randomIndex := byte(r.Intn(len(characterSet) - 1))
		bytes[i] = characterSet[randomIndex] // assign random character
	}
	return bytes
}

// Checks whether the password contains at least one number
func checkNumbers(char []byte) bool {
	for i := range numbers {
		for j := 0; j < len(char); j++ {
			if numbers[i] == char[j] {
				return true
			}
		}
	}
	return false
}

// Checks whether the password contains at least one lowercase letter
func checkLetters(char []byte) bool {
	for i := range lowercaseLetters {
		for j := 0; j < len(char); j++ {
			if lowercaseLetters[i] == char[j] {
				return true
			}
		}
	}
	return false
}

// Checks whether the password contains at least one special character
func checkSpecials(char []byte) bool {
	for i := range specialChars {
		for j := 0; j < len(char); j++ {
			if specialChars[i] == char[j] {
				return true
			}
		}
	}
	return false
}

// Checks whether the password contains at least one uppercase letter
func checkUpperCase(char []byte) bool {
	for i := range uppercaseLetters {
		for j := 0; j < len(char); j++ {
			if uppercaseLetters[i] == char[j] {
				return true
			}
		}
	}
	return false
}

// more to explore 1. encrypting passwd and storing it 2. Build LRU Cache for it. 3. Implement UI
