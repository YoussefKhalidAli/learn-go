package note

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Note struct {
	Author    string
	Title     string
	Msg       string
	CreatedAt time.Time
}

func New(author, title, msg string) *Note {
	return &Note{
		author,
		title,
		msg,
		time.Now(),
	}
}

func (n Note) ShowMsg() {
	println("author: ", n.Author)
	println("title: ", n.Title)
	println("message: ", n.Msg)
}

func (note Note) Save() {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, json, 0644)

	if err != nil {
		panic(err)
	}
}
