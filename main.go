package main

import (
	"fmt"
)

func maskLinks(input string) string {
	const prefix = "http://"
	const mask = "****************"

	inputBytes := []byte(input)
	prefixBytes := []byte(prefix)
	maskBytes := []byte(mask)

	var output []byte
	i := 0

	for i < len(inputBytes) {
		if len(inputBytes)-i >= len(prefixBytes) && string(inputBytes[i:i+len(prefixBytes)]) == prefix {
			output = append(output, prefixBytes...)
			output = append(output, maskBytes...)
			i += len(prefixBytes)
		} else {
			output = append(output, inputBytes[i])
			i++
		}
	}

	return string(output)
}

func main() {
	input := "Hello, its my page: http://localhodsersest123.com See you"
	output := maskLinks(input)
	fmt.Println("Input:", input)
	fmt.Println("Output:", output)
}
