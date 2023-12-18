package main

import (
	"fmt"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type rotor struct {
	wiring   string
	position int
}

var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"

func main() {
	rotorI := rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ"}
	rotorII := rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE"}
	rotorIII := rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO"}

	plugboard := map[rune]rune{
		'A': 'E', 'B': 'J', 'C': 'M', 'D': 'Z',
		// Add more plugboard connections as needed
	}

	plaintext := "HELLO"
	encryptedText := enigmaEncrypt(plaintext, plugboard, rotorI, rotorII, rotorIII)
	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Encrypted Text:", encryptedText)

	decryptedText := enigmaDecrypt(encryptedText, plugboard, rotorI, rotorII, rotorIII)
	fmt.Println("Decrypted Text:", decryptedText)
}

func enigmaEncrypt(plaintext string, plugboard map[rune]rune, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			// Pass the character through the rotors from right to left
			for i := len(rotors) - 1; i >= 0; i-- {
				char = substitute(char, rotors[i])
			}

			// Pass the character through the reflector
			char = reflector(char)

			// Pass the character through the rotors from left to right
			for _, r := range rotors {
				char = decrypt(char, r)
			}

			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			encrypted.WriteRune(char)

			// Rotate the rotors after each character encryption
			rotateRotors(rotors)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func enigmaDecrypt(encryptedText string, plugboard map[rune]rune, rotors ...rotor) string {
	encryptedText = strings.ToUpper(encryptedText)
	var decrypted strings.Builder

	for _, char := range encryptedText {
		if char >= 'A' && char <= 'Z' {
			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			// Pass the character through the rotors from right to left
			for i := len(rotors) - 1; i >= 0; i-- {
				char = substitute(char, rotors[i])
			}

			// Pass the character through the reflector
			char = reflector(char)

			// Pass the character through the rotors from left to right
			for _, r := range rotors {
				char = decrypt(char, r)
			}

			// Apply plugboard transformation
			if plug, ok := plugboard[char]; ok {
				char = plug
			}

			decrypted.WriteRune(char)

			// Rotate the rotors after each character decryption
			rotateRotors(rotors)
		} else {
			// Non-alphabetic characters are not modified
			decrypted.WriteRune(char)
		}
	}

	return decrypted.String()
}

func substitute(char rune, rotor rotor) rune {
	index := (int(char-'A')+rotor.position)%26 + 'A'
	return rune(rotor.wiring[index-'A'])
}

func decrypt(char rune, rotor rotor) rune {
	index := strings.IndexRune(rotor.wiring, char) - rotor.position
	if index < 0 {
		index += 26
	}
	return rune(alphabet[index])
}

func reflector(char rune) rune {
	index := strings.IndexRune(reflectorB, char)
	return rune(alphabet[index])
}

func rotateRotors(rotors []rotor) {
	rotors[0].position = (rotors[0].position + 1) % 26
	if rotors[0].position == 0 {
		rotors[1].position = (rotors[1].position + 1) % 26
		if rotors[1].position == 0 {
			rotors[2].position = (rotors[2].position + 1) % 26
		}
	}
}
