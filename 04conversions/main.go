package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Rate my website out of 5")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	fmt.Println("Thank you for rating my website")

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}

}
