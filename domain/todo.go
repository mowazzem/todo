package domain

import "errors"

type Id string

type Todo struct {
	ID          Id
	Name        string
	Description string
}

func (id Id) String() string {
	return string(id)
}

func NewId(id string) (*Id, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if len(id) > 20 {
		return nil, errors.New("id cannot be more than 20 char")
	}

	iD := Id(id)
	return &iD, nil
}

func NewTodo(id string, name, description string) (Todo, error) {
	iD, err := NewId(id)
	if err != nil {
		return Todo{}, err
	}
	todo := Todo{
		ID:          *iD,
		Name:        name,
		Description: description,
	}
	return todo, nil
}
