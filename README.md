# Greeting in go

This  package provides a simple way to get personalized greetings in Go.

## Installation

Run the following command to install the package:

```bash
go get -u github.com/chriscarsam/greetings
```

## Use
Here is an example of how to use the package in your code:

```go

import(
    "fmt"
    "github.com/chriscarsam/greetings"
)

func main() {
    message, err := greetings.Hello("Charly")

    if err != nil {
        fmt.Println("An error occurred: ", err)
        return
    }

    fmt.Println(message)
}
```
This example imports the github.com/chriscarsam/greetings package and calls the custom Hello greeting function. If an error occurs, an error message is printed.