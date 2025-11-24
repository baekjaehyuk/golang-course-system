package dto

import (
	"golang-course-registration/model"
)

type CreateLectureRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Capacity  int       `json:"capacity"`
	Credit    int       `json:"credit"`
	Day       model.Day `json:"day"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
}

type LectureResponse struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Capacity          int    `json:"capacity"`
	CurrentEnrollment int    `json:"current_enrollment,omitempty"`
	Credit            int    `json:"credit"`
	Day               string `json:"day"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
}

func NewLectureResponse(lecture model.Lecture) LectureResponse {
	return LectureResponse{
		ID:                lecture.ID,
		Name:              lecture.Name,
		Capacity:          lecture.Capacity,
		CurrentEnrollment: lecture.CurrentEnrollment,
		Credit:            lecture.Credit,
		Day:               lecture.Day.ToKorean(),
		StartTime:         lecture.StartTime,
		EndTime:           lecture.EndTime,
	}
}
