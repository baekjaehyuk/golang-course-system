package model

import (
	"errors"
	"golang-course-registration/common/constants"
	"golang-course-registration/common/exception"
)

type Student struct {
	ID int `json:"id"`
}

func NewStudent(id int) (*Student, error) {
	if id < constants.StudentIdMin || id > constants.StudentIdMax {
		return nil, errors.New(exception.ErrStudentIDInvalid)
	}

	return &Student{ID: id}, nil
}
