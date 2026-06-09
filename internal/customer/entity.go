package customer

import "time"

type Customer struct {
	ID        string     `gorm:"primaryKey;index"`
	FirstName string     `gorm:"size:100;not null"`
	LastName  string     `gorm:"size:100;not null"`
	Email     string     `gorm:"size:100;not null;unique"`
	Language  string     `gorm:"size:10;not null"`
	BirthDate *time.Time `gorm:"type:date"`
	Status    string     `gorm:"size:20;not null"`
}
