package requests

import "errors"

// StudentCreateRequest - req body struct
type StudentCreateRequest struct {
	RollNumber int    `json:"rollNumber"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	BirthDate  string `json:"birthDate"`
	Branch     string `json:"branch"`
}

// Validate - valudate req data
func (student *StudentCreateRequest) Validate() error {
	if student.RollNumber == 0 {
		return errors.New("Roll number not valid or present")
	}

	if student.FirstName == "" {
		return errors.New("FirstName not valid or present")
	}

	if student.LastName == "" {
		return errors.New("LastName not valid or present")
	}

	if student.BirthDate == "" {
		return errors.New("BirthDate not valid or present")
	}

	if student.Branch == "" {
		return errors.New("Branch not valid or present")
	}

	return nil
}
