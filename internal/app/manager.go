package app

import (
	"bufio"
	"fmt"
	"os"
)

func Manager() {
	fmt.Println("Go CLI Task Manager")
	fmt.Println("---------------------")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You entered:", text)
		if text == "exit" {
			break
		}
	}

	fmt.Println("---------------------")
	fmt.Println("Goodbye!")
}
