package service

import (
	todo "github.com/risqiboyevbobur/todo_app.git"
	"github.com/risqiboyevbobur/todo_app.git/pkg/repository"
)
type Authorization interface{
	CreateUser(user todo.User)(int, error)
}
type TodoList interface{

}
type TodoItem interface{

}
type Service struct{
	Authorization
	TodoList
	TodoItem	
}
func NewService(repos *repository.Repository)*Service  {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}