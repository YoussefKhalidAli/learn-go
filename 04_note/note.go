package main

import (
	"bufio"
	"os"
	"strings"

	"learn.com/note/note"
)

func main() {

	author := goodScan("What's your name? ")

	title := goodScan("Note title: ")

	msg := goodScan("What's on your mind? ")

	myNote := note.New(author, title, msg)

	myNote.ShowMsg()

	myNote.Save()

}

func goodScan(prompt string) string {
	print(prompt)
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(msg)
}
