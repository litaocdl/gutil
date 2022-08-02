package main

import "fmt"

func main() {
	// type check
	var i interface{} = "hello"

	// check if interface i holds concrete type string, if yes, convert, if no panic
	t := i.(string)
	fmt.Println(t)

	s, ok := i.(string)
	fmt.Println(s, ok)

	s1, ok := i.(float64)
	fmt.Println(s1, ok)

	switch o := i.(type) {
	case string:
		fmt.Printf("this is string %v", o)
	case float64:
		fmt.Printf("this is float64 %v", o)
	case int:
		fmt.Printf("this is int %v", o)
	default:
		fmt.Printf("this is default %v", o)

	}
}
