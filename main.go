
package main
import "fmt"

func main() {
    var name string // variable to store user input

    fmt.Println("Enter your name:")
    fmt.Scanln(&name) // reads input from user

    fmt.Println("Hello,", name, "! Welcome to Go 🚀") // greeting message
    fmt.Println("Keep pushing forward — you are capable of amazing things! 💪🔥") // motivational message
}