package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

// User is ...
type User struct {
	ID    string `json:"id"    valid:"uuid"`
	Name  string `json:"name"  valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

// NewUser is ...
func NewUser(id string, name string, email string) (*User, error) {
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	user.ID = uuid.NewV4().String()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
