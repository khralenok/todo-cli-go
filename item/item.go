package item

import (
	"errors"
	"fmt"
	"strconv"
)
	
type Item struct {
	Title string `json:"title"`
	Status bool `json:"status"`
}

func New(title, status string)(Item, error){
	convertedStatus, err := strconv.ParseBool(status)

	if err != nil {
		return Item{}, errors.New("Wrong status format")
	}

	if title == "" {
		return Item{}, errors.New("Title is empty")
	}

	return Item{
		Title: title,
		Status: convertedStatus,
	}, nil
}

func (item Item) Display(){
	fmt.Print(item.Title)
	fmt.Println(" ", item.DisplayStatus())
}

func (item *Item) SwitchStatus(){
	if !item.Status {
		item.Status = true
		return
	}
	item.Status = false
}

func (item Item) DisplayStatus()string{
	if item.Status {
		return "(+)"
	}
	return "(-)"
}
