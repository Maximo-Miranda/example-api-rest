package user

import (
	"github.com/jinzhu/gorm"


	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/Maximo-Miranda/example-api-rest/tools"
)

// validateCreateClientRequest ...
func validateCreateClientRequest(req *User) error {

	if err := validation.ValidateStruct(req,
		// Full Namme cannot be empty
		validation.Field(&req.FullName, validation.Required),
		// DNI cannot be empty
		validation.Field(&req.Dni, validation.Required),
	); err != nil {

		return err
	}

	err := validation.Errors{
		"dni": validation.Validate(req, validation.By(validationCheckDniIfExist)),
	}.Filter()
	if err != nil {
		return err
	}

	var errorRules error

	resultChecked, err := tools.CheckedIfKeyExist(req, "date_of_birth")
	if err != nil {
		return err
	}

	if resultChecked {
		// Name cannot be empty, and the length must between 3 and 50
		errorRules = validation.ValidateStruct(req, validation.Field(&req.DateOfBirth, validation.Date("2006-01-02")))
		return errorRules
	}

	return nil

}

// validationCheckAccountIfExist ...
func validationCheckDniIfExist(data interface{}) error {

	m := User{}
	user := data.(*User)

	m.Dni = user.Dni

	_, err := m.FirstByQuery()
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return err
	}

	return errors.New("DNI is already exist.")
}