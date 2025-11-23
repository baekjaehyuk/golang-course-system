package model

import (
	"errors"
	"golang-course-registration/common/exception"
	"time"
)

const (
	LectureIDMinLength = 1000
	LectureIDMaxLength = 9999
)

type Lecture struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Capacity          int    `json:"capacity"`
	CurrentEnrollment int    `json:"current_enrollment"`
	Credit            int    `json:"credit"`
	Day               Day    `json:"day"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
}

func NewLecture(id int, name string, capacity int, credit int, day Day, startTime, endTime string) (*Lecture, error) {
	if id < LectureIDMinLength || id > LectureIDMaxLength {
		return nil, errors.New(exception.ErrLectureIDInvalid)
	}

	if len(name) < 2 || len(name) > 20 {
		return nil, errors.New(exception.ErrLectureNameRequired)
	}

	if capacity < 1 || capacity > 30 {
		return nil, errors.New(exception.ErrLectureCapacityInvalid)
	}

	if credit < 1 || credit > 6 {
		return nil, errors.New(exception.ErrLectureCreditInvalid)
	}

	if day.ToKorean() == "undefined" {
		return nil, errors.New(exception.ErrLectureDayRequired)
	}

	start, _ := time.Parse("15:04", startTime)
	end, _ := time.Parse("15:04", endTime)
	if start.IsZero() || end.IsZero() {
		return nil, errors.New(exception.ErrLectureTimeRequired)
	}
	if !end.After(start) {
		return nil, errors.New(exception.ErrLectureTimeOrderInvalid)
	}

	return &Lecture{
		ID:                id,
		Name:              name,
		Capacity:          capacity,
		CurrentEnrollment: 0,
		Credit:            credit,
		Day:               day,
		StartTime:         startTime,
		EndTime:           endTime,
	}, nil
}

func (l *Lecture) IsFull() bool {
	return l.CurrentEnrollment >= l.Capacity
}

// IncrementEnrollment 현재 수강 인원 증가
func (l *Lecture) IncrementEnrollment() {
	l.CurrentEnrollment++
}

// DecrementEnrollment 현재 수강 인원 감소
func (l *Lecture) DecrementEnrollment() {
	if l.CurrentEnrollment > 0 {
		l.CurrentEnrollment--
	}
}
