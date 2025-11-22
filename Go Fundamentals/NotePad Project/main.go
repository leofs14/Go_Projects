package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/Note"
)

func main() {
	
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded.")
}

func getNoteData () (string, string) {
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