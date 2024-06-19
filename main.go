package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func maskURLs(input string) string {
	result := []byte(input)
	urlPrefix := []byte("http://")
	for i := 0; i < len(result); i++ {
		match := true
		for j := 0; j < len(urlPrefix); j++ {
			if i+j >= len(result) || result[i+j] != urlPrefix[j] {
				match = false
				break
			}
		}
		if match {
			j := i + len(urlPrefix)
			for ; j < len(result); j++ {
				if result[j] == ' ' {
					break
				}
				result[j] = '*'
			}
			i = j
		}
	}
	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	text := scanner.Text()
	//fmt.Println(text)

	data := []byte(text)
	maskedData := maskURLs(string(data))

	fmt.Println(maskedData)

}
func TestMaskURLs(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Visit our website at http://www.example.com for more information.", "Visit our website at http://************* for more information."},
		{"Check out https://www.google.com for search.", "Check out https://************* for search."},
		{"No URLs here.", "No URLs here."},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output := maskURLs(test.input)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected: %s, but got: %s", test.expected, output)
			}
		})
	}
}
