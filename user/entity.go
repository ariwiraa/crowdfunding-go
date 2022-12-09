package user

import (
	"time"
)

type User struct {
	ID uint32	`gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Occupation string `gorm:"type:varchar(100);not null" json:"occupation"`
	Email string `gorm:"type:varchar(100);not null:unique" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
	Avatar string `gorm:"type:varchar(255)" json:"avatar"`
	Role string `gorm:"type:varchar(10);not null" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

