package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jotaro-biz/go-remote-debug/service"
)

func main() {
	fmt.Println("start todo app.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1: create task")
		fmt.Println("2: show todo list")
		fmt.Println("3: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			fmt.Println("Please write title for todo task.")
			fmt.Print("please enter >")

			scanner.Scan()
			in = scanner.Text()
			if in == "" {
				fmt.Println("Invalid command.")
				continue
			} else {
				service.CreateTask(in)
			}
			continue
		case "2":
			fmt.Println("Here are your todo list.")
			service.ShowList()
			continue
		case "3":
			fmt.Println("bye.")
			goto M
		default:
			fmt.Println("Invalid command.")
			continue
		}
	}
M:
}
