package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Library of Alexandria")
	fmt.Println("------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Select an action by pressing \n1. For Borrowing a Book \n2. For Returning a Book \n3. To Exit")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
