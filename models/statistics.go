package models

import (
	"html/template"
	"strconv"
	"time"
)

type Statistics struct {
	ID         uint `gorm:"primary_key,column:cpu"`
	CPU        uint `gorm:"column:cpu"`
	Likes      uint `gorm:"column:likes"`
	Sales      uint `gorm:"column:sales"`
	NewMembers uint `gorm:"column:new_members"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func FirstStatics() *Statistics {
	s := new(Statistics)
	orm.First(s)
	return s
}

func (s *Statistics) CPUTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.CPU)))
}

func (s *Statistics) LikesTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.Likes)))
}

func (s *Statistics) SalesTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.Sales)))
}

func (s *Statistics) NewMembersTmpl() template.HTML {
	return template.HTML(strconv.Itoa(int(s.NewMembers)))
}
