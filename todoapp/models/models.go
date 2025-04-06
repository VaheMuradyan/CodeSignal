package models

import "time"

type Todo struct {
	ID           int       `json:"id"`
	Title        string    `json:"title" binding:"required,maxlength"`
	Completed    bool      `json:"completed"`
	CreationTime time.Time `json:"creationTime" binding:"notpast"`
	ImagePath    string    `json:"imagePath"`
}
