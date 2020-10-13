package greetings

import (
    "errors"
    "fmt"
)


// Hello returns a greeting for the named person.
// func {Function Name}({Patameter Name} {Parameter Type}) {Return Type}
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name") // returns an error with message
    }

    // If a name was received, return a value that embeds the name in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // declaring and initializing a variable in one line
    return message, nil // nil meaning no error
}
