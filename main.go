package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	bytes := []byte(text)

	usr := bytes[:]
	for _, b := range bytes {
		fmt.Printf("%v", b)
		letter := string(b)
		fmt.Printf("%v\n", letter) // Выводим значение каждого байта
	}
	fmt.Println()
	fmt.Println(string(usr))

	found := false
	for b := 0; b < len(bytes)-1; b++ {

		if bytes[b] == 47 && bytes[b-1] == 47 && bytes[b-2] == 58 && bytes[b-3] == 112 && bytes[b-4] == 116 && bytes[b-5] == 116 && bytes[b-6] == 104 {
			found = true
			continue
		}
		if found {
			if bytes[b] == 32 {
				break
			}
			bytes[b] = byte('*')

		}
	}
	fmt.Printf("Маскированный байтовый срез: %v\n", string(usr))
}
