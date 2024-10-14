package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Println("Enter add, remove, complete, list or exit")
		userInput := scanner.Scan()
		if !userInput {
			break
		}
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Println("Enter task to add:")
			scanner.Scan()
			todo := scanner.Text()

			// Split task into multiple items if they are separated by spaces
   			new := strings.Split(todo, " ")

			for _, item := range new {
				tasks = append(tasks, item)
			}
            // tasks = append(tasks, todo)

		case "remove":
			fmt.Println("Enter task number(s) to remove:")
			scanner.Scan()
			userInput := scanner.Text()
			// if !isNumeric(userInput) {
			// 	fmt.Println("Invalid input. Please enter numeric values.")
            //     break
            // }
			new := strings.Split(userInput, " ")
			if len(new) < 2 {
				taskNumber, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Couldn't parse task number")
				}
	
				if taskNumber < 1 || taskNumber > len(tasks) {
					fmt.Println("Invalid task number, must start from 1")
				}
				tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
			} else {
				for ind, taskNumber := range new {
                    taskNumber, err := strconv.Atoi(taskNumber)
                    if err!= nil {
                        fmt.Println("Couldn't parse task number")
                        continue
                    }
                    if taskNumber < 1 || taskNumber > len(tasks) {
                        fmt.Println("Invalid task number, must start from 1")
                        continue
                    }
					if ind == 0 {
						tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)
					} else {
						tasks = append(tasks[:taskNumber-2], tasks[taskNumber-1:]...)
					}
                }
			}

		case "complete":
			fmt.Println("Enter task numbers to mark as completed.")
			scanner.Scan()
			userInput := scanner.Text()
			new := strings.Split(userInput, " ")
			for _, taskNumber := range new {
				taskNumber, err := strconv.Atoi(taskNumber)
                if err!= nil {
                    fmt.Println("Couldn't parse task number")
                    continue
                }
                if taskNumber < 1 || taskNumber > len(tasks) {
                    fmt.Println("Invalid task number, must start from 1")
                    continue
                }
				fmt.Printf("Task '%s' marked complete.\n", tasks[taskNumber-1])
				tasks[taskNumber-1] = fmt.Sprintf("(Task '%s' completed)\n", tasks[taskNumber-1])
				
			}
		case "list":
			if len(tasks) == 0 {
				fmt.Println("You don't have any tasks yet.")
			} else {
				fmt.Println("Your tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "exit":
			fmt.Println("Exiting...")
            return

		default:
			fmt.Println("Invalid command. Please try again.")
	
		}
	}
}


func isNumeric(s string) bool {
	re := regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)
	return re.MatchString(s)
}