// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct{
    name string
    age int
    email string
}

// main is the entry point for the application.
func main() {
    // Declare variable of type user and init using a struct literal.
    felipe := user{
        name: "Felipe",
        age: 32,
        email: "felipef.silva@gmail.com",
    }

    // Display the field values.
    fmt.Println("name", felipe.name)
    fmt.Println("age", felipe.age)
    fmt.Println("email", felipe.email)

    // Declare a variable using an anonymous struct.
    max := struct{
        name string
        age int
        email string
    }{
        name: "Maxwell",
        age: 31,
        email: "maxwell.dasilva@gmail.com",
    }
    
    // Display the field values.
    fmt.Println("name", max.name)
    fmt.Println("age", max.age)
    fmt.Println("email", max.email)
}
