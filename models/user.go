package models

// User represents a user entity
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Users is the in-memory data store (hard-coded mock data)
var Users = []User{
	{ID: 1, Name: "Nguyen Van A", Email: "nguyenvana@example.com", Phone: "0901234567"},
	{ID: 2, Name: "Tran Thi B", Email: "tranthib@example.com", Phone: "0912345678"},
	{ID: 3, Name: "Le Van C", Email: "levanc@example.com", Phone: "0923456789"},
}
