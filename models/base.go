package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	conn db.Connection
	orm  *gorm.DB
	err error
)

func SetConn(c db.Connection) {
	conn = c
}

func Init() {
	orm, err = gorm.Open("sqlite", conn.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
}
