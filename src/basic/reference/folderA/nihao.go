package greetings

import "fmt"

// Hello returns a greeting for the named person.
func CallMeNihao(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("NiHao, %v. Welcome nihao!", name)
	return message
}
