// Code contained in this file is used to solve challenges from: https://cryptopals.com/sets/1
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("Solutions for Set 1:")
	fmt.Println("\n1.1:", convertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println("1.2:", fixedXorDecrypt("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
	answer, score, err := singleXorCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err == nil {
		fmt.Println("1.3:", string(answer), score)
	} else {
		fmt.Println("1.3:", err)
	}

}

// Decodes a hex string into a array of bytes.
func convertToHex(hexStr string) []byte {
	decodedHex, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	return decodedHex
}

// Use most frequent letters in english plain text for scoring.
func getCharWeight(char byte) int {
	wm := map[byte]int{
		byte('U'): 2,
		byte('u'): 2,
		byte('L'): 3,
		byte('l'): 3,
		byte('D'): 4,
		byte('d'): 4,
		byte('R'): 5,
		byte('r'): 5,
		byte('H'): 6,
		byte('h'): 6,
		byte('S'): 7,
		byte('s'): 7,
		byte(' '): 8,
		byte('N'): 9,
		byte('n'): 9,
		byte('I'): 10,
		byte('i'): 10,
		byte('O'): 11,
		byte('o'): 11,
		byte('A'): 12,
		byte('a'): 12,
		byte('T'): 13,
		byte('t'): 13,
		byte('E'): 14,
		byte('e'): 14,
	}
	return wm[char]
}

// Converts a hex string into it's base64 equivalent.
func convertHexToBase64(hexStr string) string {
	bytes := convertToHex(hexStr)
	baseStr := base64.StdEncoding.EncodeToString(bytes)
	return baseStr
}

// Hex decodes two strings and XOR's them against each other.
func fixedXorDecrypt(bufferA string, bufferB string) string {
	decodedA := convertToHex(bufferA)
	decodedB := convertToHex(bufferB)
	if len(decodedA) != len(decodedB) {
		fmt.Println("buffers must have equal length")
		return ""
	}

	result := make([]byte, len(decodedA))
	for i := 0; i < len(decodedA); i++ {
		result[i] = decodedA[i] ^ decodedB[i]
	}

	return hex.EncodeToString(result)
}

// Finds the decription key to decrypt a string XOR'd against a single character.
func singleXorCipher(codedStr string) ([]byte, int, error) {
	bytes := convertToHex(codedStr)
	var answer []byte
	var score int
	for i := 0; i < 256; i++ {
		r := make([]byte, len(bytes))
		var s int
		for j := 0; j < len(bytes); j++ {
			c := bytes[j] ^ byte(i)
			s += getCharWeight(c)
			r[j] = c
		}
		if s > score {
			answer = r
			score = s
		}
		s = 0
	}
	return answer, score, nil
}
