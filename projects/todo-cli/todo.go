package main

import (
	"fmt"
)

type Todo struct {
	title  string
	status string
	id     int
}

func main() {
	var tasks []Todo
	var userChoice int

	for {
		fmt.Println("1. Add Todo 2. View Todos 3. Delete Todo 4. Exit application")
		fmt.Printf("Enter your choice : ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			var taskTitle string
			var todoId int
			fmt.Printf("Enter the task title : ")
			fmt.Scan(&taskTitle)
			fmt.Printf("Enter todo id : ")
			fmt.Scan(&todoId)
			task := Todo{title: taskTitle, status: "Pending", id: todoId}
			tasks = append(tasks, task)
			fmt.Println("Task Added successfully")
			break
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks are founded to view")
				break
			}
			fmt.Printf("Tasks : \n")
			for i, val := range tasks {
				fmt.Printf("%d : %s %s\n", i+1, val.title, val.status)
			}
			break
		case 3:
			var todoId int
			tasklen := len(tasks)
			counter := 0
			newTodos := []Todo{}
			fmt.Printf("Enter todoid to delete : ")
			fmt.Scan(&todoId)

			for _, val := range tasks {
				counter++
				if val.id != todoId {
					newTodos = append(newTodos, val)
				}
				if counter != tasklen {
					fmt.Println("Invalid todo id founded")
				}
			}
			tasks = append(newTodos)
			fmt.Println("Todo deleted successfully!")
			break
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid User choice")
		}
	}
}
