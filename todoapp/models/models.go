package models

import "time"

type Todo struct {
	ID           int       `json:"id"`
	Title        string    `json:"title" binding:"required"`
	Completed    bool      `json:"completed"`
	CreationTime time.Time `json:"creationTime" binding:"notpast"`
}
