package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to TODO List Cli by ForzzMinay")
	var ToDo []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter your command(create, read, update, delete, exit):")
		var command string
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error", err)
			continue
		}

		command = strings.TrimSpace(strings.ToLower(command))

		if command == "create" {
			fmt.Println("Enter your task:")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			ToDo = append(ToDo, input)
			fmt.Println("Task added successfully")
		} else if command == "read" {
			if len(ToDo) == 0 {
				fmt.Println("List is empty")
			}
			for i, task := range ToDo {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		} else if command == "exit" {
			fmt.Println("Have a good day!")
			break
		} else if command == "delete" {
			if len(ToDo) == 0 {
				fmt.Println("List is empty")
				continue
			}
			for i, task := range ToDo {
				fmt.Printf("%d. %s\n", i+1, task)
			}
			fmt.Println("Enter the task number")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			index := num - 1
			if err != nil || index < 0 || index >= len(ToDo) {
				fmt.Println("Numbering error", err)
				continue
			}

			ToDo = append(ToDo[:index], ToDo[index+1:]...)
			fmt.Println("Task deleted successfully")
		} else if command == "update" {
			if len(ToDo) == 0 {
				fmt.Println("List is empty")
				continue
			}
			for i, task := range ToDo {
				fmt.Printf("%d. %s\n", i+1, task)
			}
			fmt.Println("Enter the task number")
			input, err := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			index := num - 1
			if err != nil || index < 0 || index >= len(ToDo) {
				fmt.Println("Numbering error", err)
				continue
			}
			fmt.Println("What`s new task?")
			inputtask, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input failed. Error:", err)
			}
			inputtask = strings.TrimSpace(inputtask)
			ToDo[index] = inputtask
			fmt.Println("Task updated successfully")
		} else {
			fmt.Println("Invalid command")
		}
	}
}
