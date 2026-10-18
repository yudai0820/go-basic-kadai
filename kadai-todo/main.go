package main

import "fmt"

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func (t *Todo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.Id)
	t.Completed = true
}

func printTodos(todos []Todo) {
	for _, todo := range todos {
		status := "未完了"
		if todo.Completed {
			status = "完了"
		}
		fmt.Printf("[%s]（ID：%d）%s\n", status, todo.Id, todo.Title)
	}
}

func main() {
	todos := []Todo{
		{Id: 1, Title: "学習計画", Completed: false},
		{Id: 2, Title: "環境構築", Completed: false},
		{Id: 3, Title: "基礎文法", Completed: false},
	}

	fmt.Println("--- ToDoリスト（初期状態） ---")
	printTodos(todos)

	todos[0].Complete()
	todos[1].Complete()

	fmt.Println("--- ToDoリスト（最終状態） ---")
	printTodos(todos)
}
