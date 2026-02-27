package models

// User represents a user entity mapped to "users" table in Supabase
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Email string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Phone string `json:"phone" gorm:"type:varchar(20)"`
}
