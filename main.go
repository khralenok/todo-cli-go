package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/todo_app/list"
)

//To do:
//1. Distribute functions to corresponding separate modules
//2. Handle errors more precice

func main(){
	var thisList = list.New()
	
	fmt.Println("Welcome to Todo.app")

	for {
		thisList.OutputList()
		if outputMenu(&thisList) {
			return
		}
	}
}


func outputMenu(todoList *list.List) bool{
	fmt.Println("MENU:")
	fmt.Println("--------------------")
	fmt.Println("1. Add new item to the list")
	fmt.Println("2. Switch item status")
	fmt.Println("3. Remove item from the list")
	fmt.Println("4. Exit the app")
	fmt.Println("--------------------")
	return router(getMenuItemChoice("Type menu item number: "), todoList)
}

func router(userInput int, todoList *list.List) bool {
	switch userInput{
	case 1:
		err := todoList.AddItem(getUserInput("Type what you want to do: "))
		if err != nil{
			return true
		}
		return false
	case 2:
		index := getMenuItemChoice("What item you want to switch status for: ")-1
		todoList.Items[index].SwitchStatus()
		return false
	case 3:
		index := getMenuItemChoice("What item you want to remove: ")-1
		todoList.RemoveItem(index)
		return false
	case 4:
		return true
	default: 
		fmt.Println("Such feature doesn't exist. Check available options in menu")
		return false
	}
}

func getMenuItemChoice(prompt string) int{
	var userInput int
	fmt.Print(prompt)
	fmt.Scanln(&userInput)
	return userInput
}

func getUserInput(prompt string)string{
	fmt.Print(prompt)
	reader:=bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err!=nil{
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	
	return text
}