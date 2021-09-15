package task5

import (
	"time"
)

type Ad struct {
	Id 			int 		`json:"id" db:"id" `
	Title 		string 		`json:"title" binding:"required" db:"title" validate:"required,lte=200"`
	Price       int  		`json:"price" db:"price"`
	MainPhoto   string  	`json:"main_photo" db:"main_photo"`
	Photo1      string  	`json:"photo1" db:"photo1"`
	Photo2      string  	`json:"photo2" db:"photo2"`
	Photo3      string  	`json:"photo3" db:"photo3"`
	Description string  	`json:"description" db:"description" validate:"required,lte=1000"`
	Date 		time.Time 	`json:"date" db:"date"`
}
