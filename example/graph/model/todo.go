package model

import "github.com/cloudmatelabs/gorm-gqlgen-relay/relay"

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID     int     `json:"id" gorm:"primaryKey,autoIncrement"`
	Text   string  `json:"text"`
	Done   bool    `json:"done"`
	UserID *string `json:"userId,omitempty"`
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TodoEdge = relay.Edge[Todo]
type TodoConnection = relay.Connection[Todo]
