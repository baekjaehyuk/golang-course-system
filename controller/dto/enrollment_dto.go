package dto

import (
	"golang-course-registration/model"
)

type EnrollRequest struct {
	StudentID int `json:"student_id"`
	LectureID int `json:"lecture_id"`
}

type EnrollmentResponse struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	LectureID int `json:"lecture_id"`
}

func NewEnrollmentResponse(enrollment model.Enrollment) EnrollmentResponse {
	return EnrollmentResponse{
		ID:        enrollment.ID,
		StudentID: enrollment.StudentID,
		LectureID: enrollment.LectureID,
	}
}
