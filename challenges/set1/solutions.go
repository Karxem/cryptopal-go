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
}

func convertToHex(hexStr string) []byte {
	decodedHex, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	return decodedHex
}
func convertHexToBase64(hexStr string) string {
	bytes := convertToHex(hexStr)
	baseStr := base64.StdEncoding.EncodeToString(bytes)
	return baseStr
}

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
