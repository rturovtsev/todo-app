package main

import (
	"github.com/rturovtsev/todo-app"
	"github.com/rturovtsev/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}
}
