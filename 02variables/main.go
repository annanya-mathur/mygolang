package main

import "fmt"

func main() {
	var name string = "ABC"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var num uint = 256
	fmt.Println(num)
	fmt.Printf("Variable is of type: %T \n", num)

	//implicit type
	var website = "https://github.com/annanya-mathur"
	fmt.Println(website)

	number := 3000.0
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

}
