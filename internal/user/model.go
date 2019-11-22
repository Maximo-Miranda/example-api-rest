package user

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/Maximo-Miranda/example-api-rest/tools"
)

// User ...
type User struct {
	ID         		uuid.UUID       `gorm:"type:uuid;primary_key;" json:"id"`
	FullName		string			`json:"full_name"`
	Dni				string			`json:"dni"`
	DateOfBirth		*string			`json:"date_of_birth"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}


// BeforeCreate will set a UUID rather than numeric ID.
func (base *User) BeforeCreate(scope *gorm.Scope) error {

	return scope.SetColumn("ID", uuid.NewV4())
}

// Save ...
func (m *User) Save() (*User, error) {

	conn, err := tools.Connect()
	if err != nil {
		return m, err
	}
	defer conn.Close()

	result := conn.Create(m)
	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}
