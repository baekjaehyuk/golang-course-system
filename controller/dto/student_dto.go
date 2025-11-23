package dto

import (
	"golang-course-registration/model"
)

type CreateStudentRequest struct {
	ID int `json:"id"`
}

type StudentResponse struct {
	ID int `json:"id"`
}

func NewStudentResponse(student model.Student) StudentResponse {
	return StudentResponse{
		ID: student.ID,
	}
}
