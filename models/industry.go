package models

import (
	"github.com/jinzhu/gorm"
)

//Industry struct holds struct data
type Industry struct {
	gorm.Model
	Industry string `gorm:"type:varchar(100)" json:"industry"`
}
