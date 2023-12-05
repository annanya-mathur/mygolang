package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Rate my website out of 5")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	fmt.Println("Thank you for rating my website")
	fmt.Println(input, err)
}
