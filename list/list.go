package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/todo_app/filemanager"
	"example.com/todo_app/item"
)

type List struct{
	Items []item.Item `json:"todo_list"`
}

func New() List {
	listFromFile, err := filemanager.ReadFile("todo.json")
	if err != nil{
		var newList []item.Item
		return List{
			Items: newList,
		}
	}

	return List{
		Items: listFromFile,
	}
}

func (list List) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(list)

	if err != nil{
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (list *List) AddItem(itemName string) error{
	newNote, err := item.New(itemName, "false")
	if err!=nil{
		return errors.New("New item creation was failed")
	}
	list.Items = append(list.Items, newNote)
	list.Save()
	return nil
}



func (list *List) RemoveItem(itemIndex int) {
	var newTodo []item.Item
	for index, value := range list.Items{
		if index != itemIndex{
			newTodo = append(newTodo, value)
		}
	}

	list.Items = newTodo
	list.Save()
}

func (list List) OutputList(){
	fmt.Println("--------------------")
	fmt.Println("YOUR TODO LIST:")
	fmt.Println("--------------------")
	for index, value := range list.Items {
		fmt.Printf("%v. ", index+1)
		fmt.Print(value.Title)
		fmt.Println(" ", value.DisplayStatus())
	}
	fmt.Println("--------------------")
	fmt.Println("--------------------")
}