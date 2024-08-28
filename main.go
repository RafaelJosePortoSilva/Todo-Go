package main

import (
	_ "bufio"
	"fmt"
	_ "fmt"
	_ "os"
	_ "strconv"
	_ "strings"
)

// Estrutura da task
type Task struct {
	Text      string
	Completed bool
}

func main() {

	// Estrutura de dados
	tasks := []Task{}

	// Loop principal da aplicação
	for {

		showMenu()
		option := getUserInput("Enter your choice: ")

		switch option {

		case "1":
			showTasks(tasks)
		case "2":
			addTask(&tasks)
		case "3":
			marktaskCompleted(&tasks)
		case "4":
			saveTasksToFile(&tasks)
		case "5":
			fmt.Println("Exiting from ToDo-App...")
			return
		default:
			fmt.Println("Invalid choice. Please, try again.")

		}

	}
}

func showMenu() {

	fmt.Println("\nMenu:")
	fmt.Println("[1] Show Tasks")
	fmt.Println("[2] Add Task")
	fmt.Println("[3] Mark Task as Completed")
	fmt.Println("[4] Save Tasks to File")
	fmt.Println("[5] Exit")

}
