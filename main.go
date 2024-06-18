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
	Find := []byte("http://")
	Input := []byte(text)

	for i := 0; i < len(Input)-len(Find)+1; i++ {
		if string(Input[i:i+len(Find)]) == "http://" {

			for j := i + len(Find); j < len(Input); j++ {
				if Input[j] == ' ' {

					for k := i + len(Find); k < j; k++ {
						Input[k] = '*'
					}
					break
				}
			}
		}
	}

	fmt.Println("Результат:", string(Input)) // Вывод результата
}
