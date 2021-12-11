package service

import "github.com/KrissLa/todo_app_go/pkg/repository"

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
