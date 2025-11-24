package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/Note"
	"example.com/notepad/Todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver // Embedding and interface with another
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = dataDisplayer(todo)

	if err != nil {
		return
	}

	err = dataDisplayer(userNote)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) { // It can be used with any value
	switch value.(type){
	case int:
		fmt.Println("Integer", value)
	case float64:
		fmt.Println("Float64", value)
	case string:
		fmt.Println(value)
	}
}

func dataDisplayer(data outputtable) error {
	data.Display()
	saveData(data)
	return saveData(data)
}

func saveData(data saver) error{
		err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin) // STDin means it will read from your command line.

	text, err := reader.ReadString('\n') // Use single quote to identify a single character

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") //Windows use /r/n to linebreak.

	return text
}