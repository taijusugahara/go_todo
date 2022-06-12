package main

import (
	// "fmt"
	// "log"
	"todo_app/app/contorollers"
	// "todo_app/app/models"
	// "todo_app/config"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// log.Println("test")

	// fmt.Println(models.Db)

	//////create
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	//////get
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	/////update
	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	/////delete
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	//////Todo
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Hello")
	// user.CreateTodo("Good morning")
	// user.CreateTodo("Good afternoon")
	// user.CreateTodo("Good evening")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user, _ := models.GetUser(3)
	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "First update done"
	// t.UpdateTodo()

	// t, _ := models.GetTodo(3)
	// t.DeleteTodo()

	contorollers.StartMainServer()
}
