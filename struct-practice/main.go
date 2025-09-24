package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"kayodedev.com/prac/note"
	"kayodedev.com/prac/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todoData, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ouputData(todoData)

	if err != nil {
		return
	}
	ouputData(data)
}

func ouputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	if err != nil {
		return err
	}
	return nil
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("File counln't be saved")
		return err
	}
	fmt.Println("File saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
