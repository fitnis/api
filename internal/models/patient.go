package models

import "time"

type Patient struct {
    ID        uint      `gorm:"primaryKey"`
    FirstName string    `gorm:"size:100;not null"`
    LastName  string    `gorm:"size:100;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
