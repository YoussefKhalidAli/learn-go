package main

import (
	"bufio"
	"os"

	"learn.com/todo/todo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	print("What are you trying to accomplish? ")
	task, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	myTodo := todo.New(task)

	myTodo.Show()
	finishTask(myTodo)
	myTodo.Show()
}

func finishTask(todo todo.TodoOps) {
	todo.Finish()
}
