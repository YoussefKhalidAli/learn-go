package todo

type Todo struct {
	task string
	done bool
}

type TodoOps interface {
	Finish()
	Show()
}

func New(task string) *Todo {
	todo := Todo{
		task,
		false,
	}

	return &todo
}

func (t *Todo) Finish() {
	t.done = true
}

func (t Todo) Show() {
	println(t.task, t.done)
}
