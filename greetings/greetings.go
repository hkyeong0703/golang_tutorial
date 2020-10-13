package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)


// Hello returns a greeting for the named person.
// func {Function Name}({Patameter Name} {Parameter Type}) {Return Type}
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name") // returns an error with message
    }

    // Create a message using a random format.
    // randomFormat function that returns a randomly selected format for a greeting message. 
    //      starts with a lowercase letter, making it accessible only to code in its own package
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil

    // If a name was received, return a value that embeds the name in a greeting message.
    // message := fmt.Sprintf("Hi, %v. Welcome!", name) // declaring and initializing a variable in one line
    //  return message, nil // nil meaning no error
}

// Hellos returns a map that associates each fo the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A amp to associate names with messages.
    messages := make(map[string]string)

    // Loop through the received slice of names, calling the Hello function to get a message fot each name.
    for _, name := range names { // You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with the name.
        messages[name] = message
    }
    return messages, nil
} 

// init sets initial values for variables used in the funtion.
func init() {
    rand.Seed(time.Now().UnixNano())
}


// randomFormat rueturns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    // Go that the array underlying a slice can be dynamically sized/
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
