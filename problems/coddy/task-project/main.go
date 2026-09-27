package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func addTask(tasks []Task, taskName string, status bool) []Task {
	newTask := Task{taskName, status}
	return append(tasks, newTask)
}

func viewAllTasks(tasks []Task) {
	completedTaskCount := 0
	for _, task := range tasks {
		if task.Completed {
			fmt.Printf("[x] %s\n", task.Name)
			completedTaskCount++
		} else {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(tasks), completedTaskCount, len(tasks)-completedTaskCount)
}

func completeTask(tasks *[]Task, index int) {
	(*tasks)[index].Completed = true
}
func removeTask(tasks []Task, index int) []Task {
	updatedTasks := append(tasks[:index], tasks[index+1:]...)
	return updatedTasks
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := scanner.Text()

	scanner.Scan()
	line2 := scanner.Text()

	scanner.Scan()
	line3 := scanner.Text()

	tasks := []Task{}

	parts := strings.Split(line1, ",")
	appName := parts[0]
	userName := parts[1]
	fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Remove Task")
	fmt.Println("5. Exit")

	var newTasks []string
	if line2 != "" {
		newTasks = strings.Split(line2, ",")
	}
	for i := 0; i < len(newTasks); i++ {
		indTask := strings.Split(newTasks[i], ":")
		status := false
		if indTask[1] == "true" {
			status = true
		}
		tasks = addTask(tasks, indTask[0], status)
	}
	fmt.Printf("Current tasks: %d\n", len(tasks))

	newTodos := strings.Split(line3, ",")
	for _, todo := range newTodos {
		if todo == "exit" {
			fmt.Println("--- EXIT ---")
			break
		}
		if todo == "view" {
			fmt.Println("--- VIEW TASKS ---")
			viewAllTasks(tasks)
		}
		parts := strings.Split(todo, "|")
		todo := parts[0]
		if todo == "add" {
			tasks = addTask(tasks, parts[1], false)
			fmt.Println("--- ADD TASK ---")
			fmt.Printf("Task '%s' added!\n", parts[1])
		}
		if todo == "complete" {
			fmt.Println("--- COMPLETE TASK ---")
			index, _ := strconv.Atoi(parts[1])
			if index >= len(tasks) {
				fmt.Println("Invalid task number")
			} else {
				completeTask(&tasks, index)
				fmt.Printf("Task '%s' marked as completed!\n", tasks[index].Name)
			}

		}
		if todo == "remove" {
			fmt.Println("--- REMOVE TASK ---")
			index, _ := strconv.Atoi(parts[1])
			if index >= len(tasks) {
				fmt.Println("Invalid task number")
			} else {
				removedTaskName := tasks[index].Name
				tasks = removeTask(tasks, index)
				fmt.Printf("Task '%s' removed successfully!\n", removedTaskName)
			}
		}
	}

	fmt.Println("Final list:")
	viewAllTasks(tasks)

	fmt.Printf("Goodbye, %s!\n", userName)
}
