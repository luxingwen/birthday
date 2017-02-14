package models

import (
	"time"
)

type User struct {
	Id        int
	UserName  string
	Pwd       string
	Phone     string
	Pic       string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time
}

func InsertObj(u interface{}) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	_, err = engine.Insert(u)
	return err
}
