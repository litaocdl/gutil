/**
1. package name can different than folder name
2. but one folder can not have two pacakges
*/
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func CallMeHello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
