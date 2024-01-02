package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Structure representing a task
type Task struct {
	ID    int
	Title string
}

// Collection of tasks
var tasks []Task

// Function to add a new task
func addTask(title string) {
	taskID := len(tasks) + 1
	newTask := Task{ID: taskID, Title: title}
	tasks = append(tasks, newTask)
	fmt.Printf("Added task: %s (ID: %d)\n", title, taskID)
}

// Function to remove a task
func removeTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Removed task with ID: %d\n", taskID)
			return
		}
	}
	fmt.Printf("Task with ID: %d not found\n", taskID)
}

// Function to display the list of tasks
func listTasks() {
	fmt.Println("Task List:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Title)
	}
}

// Function to save the list of tasks to a file
func saveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		line := fmt.Sprintf("%d,%s\n", task.ID, task.Title)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()

	fmt.Printf("Saved task list to file: %s\n", filename)
	return nil
}

// Function to load the list of tasks from a file
func loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			fmt.Println("Error loading data from file.")
			return nil
		}

		taskID, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error converting task ID.")
			return nil
		}

		newTask := Task{ID: taskID, Title: fields[1]}
		tasks = append(tasks, newTask)
	}

	fmt.Printf("Loaded task list from file: %s\n", filename)
	return nil
}

func main() {
	fmt.Println("Welcome to the task management program!")

	// Load the task list from a file on startup
	loadFromFile("tasks.txt")

	for {
		fmt.Print("\nSelect an option:\n1. Add a new task\n2. Display the task list\n3. Remove a task\n4. Save the list to a file\n5. Exit\nOption: ")

		var option string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			option = scanner.Text()
		}

		switch option {
		case "1":
			fmt.Print("Enter the title of the new task: ")
			if scanner.Scan() {
				title := scanner.Text()
				addTask(title)
			}
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter the ID of the task to remove: ")
			if scanner.Scan() {
				taskIDStr := scanner.Text()
				taskID, err := strconv.Atoi(taskIDStr)
				if err != nil {
					fmt.Println("Error converting task ID.")
					break
				}
				removeTask(taskID)
			}
		case "4":
			saveToFile("tasks.txt")
		case "5":
			fmt.Println("Thank you! Goodbye.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
