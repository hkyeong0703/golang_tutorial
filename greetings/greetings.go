package greetings

import "fmt"

// Hello returns a greeting for the named person.
// func {Function Name}({Patameter Name} {Parameter Type}) {Return Type}
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // declaring and initializing a variable in one line
    return message
}
