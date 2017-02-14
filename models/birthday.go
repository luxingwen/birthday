package models

import (
	"time"
)

type BirthDay struct {
	Id              int
	Name            string
	Avatar          string
	Remark          string
	SolarCalendar   string
	ChineseCalendar string
	Own             int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func Qbirthday(own int) (r []*BirthDay, err error) {
	engine, err := GetEngine()
	if err != nil {
		return nil, err
	}
	defer engine.Clone()
	err = engine.Where("own=?", own).Find(&r)
	if err != nil {
		return nil, err
	}
	return
}
