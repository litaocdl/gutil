package main

/**
if we are use route way, suppose the under src folder the package is exported as `golang_learning`
we should reference start with golang_learning
*/
import (
	"fmt"
	greetings "golang_learning/basic/reference/folderA"
	"golang_learning/basic/reference/folderB"
)

func main() {
	fmt.Printf(" hello world\n")
	message := greetings.CallMeHello("Gladys")
	fmt.Println(message)

	message = greetings.CallMeHi("Gladys")
	fmt.Println(message)

	message = greetings.CallMeNihao("Gladys")
	fmt.Println(message)

	message = folderB.CallMeRoute("Gladys")
	fmt.Println(message)

}
