package main

import "fmt"
import "github.com/nexidian/gocliselect"
import "os"

func handleAdd() {

    var interest string
    fmt.Println("Enter your interest to add:")
    fmt.Scanln(&interest)



}

func handleView() {
}


func main() {
    fmt.Println("Simple Program to Journal your daily interests!")
	menu := gocliselect.NewMenu("Choose an Option")
    menu.AddItem("Add an Interest" , "add")
    menu.AddItem("View Interests" , "view")
    menu.AddItem("Exit" , "exit")
    choice := menu.Display()
    fmt.Printf("Choice: %s\n", choice)

    if choice=="add" {
        handleAdd()
    } else if choice=="view" {
        handleView()
    } else if choice=="exit" {
        fmt.Println("Exiting the program. Goodbye!")
    } else {
        fmt.Println("Invalid choice. Please try again.")
    }
}