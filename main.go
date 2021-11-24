package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rajbahakdanyu/lms_go/utils"
)

func main() {
	fmt.Println("Welcome to the Library of Alexandria")
	fmt.Println("------------------------------------")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select an action by pressing \n1. For Borrowing a Book \n2. For Returning a Book \n3. To Exit")
		text, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(text)
		println()

		switch choice {
		case "1":
			utils.Borrow()
		case "2":
			utils.Return()
		case "3":
			fmt.Println("Good Bye")
		default:
			fmt.Println("Please enter correct code.")
		}

		if choice == "3" {
			break
		}
	}
}
