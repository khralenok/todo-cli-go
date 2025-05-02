package filemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	"example.com/todo_app/item"
)

func ReadFile(path string) ([]item.Item, error){
	jsonData, err := os.ReadFile(path)
	
	if err != nil {
		return nil, errors.New("Some problems with reading a file")
	}

	convertedJson := make(map[string]any)

	err = json.Unmarshal(jsonData, &convertedJson)

	if err != nil{
		return nil, errors.New("Some problems with parsing json")
	}

	valueOfJson := reflect.ValueOf(convertedJson["todo_list"])
	todoList := make([]interface{}, valueOfJson.Len())
	for i:=0; i < valueOfJson.Len(); i++ {
		todoList[i] = valueOfJson.Index(i).Interface()
	}

	var newList []item.Item



	for _, value := range todoList {
		m, ok:= value.(map[string]any)
		if !ok {
			fmt.Println("Type assertation failed")
			return nil, errors.New("Type assertation failed")
		}

		newTitle := fmt.Sprintf("%v", m["title"])
		newStatus := fmt.Sprintf("%v", m["status"])
		newNote, err := item.New(newTitle, newStatus)
		if err != nil{
			return nil, errors.New("New list item creation issues")
		}

		newList = append(newList, newNote)
	}

	return newList, nil
}

func SaveFile() {

}