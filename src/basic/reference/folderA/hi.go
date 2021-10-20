package greetings

import "fmt"

// Hello returns a greeting for the named person.
func CallMeHi(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome to hi!", name)
	return message
}
