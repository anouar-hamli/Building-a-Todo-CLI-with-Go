package main

import (
	"time"
	"fmt"
)
 type Crud struct {
 	title string
	completed bool
	createdAt time.Time
	completedAt *time.Time

 }
	type Cruds []Crud
	func (c *Cruds) Add(title string) {
		Crud := Crud{
			title: title,
			completed: false,
			createdAt: time.Now(),
			completedAt: nil,
		}
		*c = append(*c, Crud)
	}
	func (c *Cruds) ValidationIndex(index int)error {
		if index < 0 || index >= len(*c) {
		error := fmt.Errorf("index %d is out of range", index)
		fmt.Println(error)
		return error
		}
		return nil
	}
	func (c *Cruds) delete(index int) error {
		t := *Cruds
		if err := t.ValidationIndex(index); err != nil {
			return err
		}
		*c = append(t[:index], t[index+1:]...)
		return nil
	}