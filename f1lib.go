package f1lib

import "fmt"

// Hello returns a greeting for the named person.
func WhosBest(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("%v is the best!", name)
	return message
}
