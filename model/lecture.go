package model

import (
	"errors"
	"golang-course-registration/common/constants"
	"golang-course-registration/common/exception"
	"time"
)

type Lectures []Lecture

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
	li, err := validateLectureId(id)
	if err != nil {
		return li, err
	}

	ln, err := validateLectureName(name)
	if err != nil {
		return ln, err
	}

	lcp, err := validateLectureCapacity(capacity)
	if err != nil {
		return lcp, err
	}

	lc, err := validateLectureCredit(credit)
	if err != nil {
		return lc, err
	}

	ld, err := validateLectureDay(day)
	if err != nil {
		return ld, err
	}

	lt, err := validateLectureTime(startTime, endTime)
	if err != nil {
		return lt, err
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

// IncrementCurrentEnrollment 현재 수강 인원 업데이트
func (l *Lecture) IncrementCurrentEnrollment() {
	l.CurrentEnrollment++
}

// DecrementCurrentEnrollment 현재 수강 인원 업데이트
func (l *Lecture) DecrementCurrentEnrollment() {
	if l.CurrentEnrollment > 0 {
		l.CurrentEnrollment--
	}
}

// HasTimeConflict 시간 중복 검증
func (l *Lecture) HasTimeConflict(other *Lecture) bool {
	if l.Day != other.Day {
		return false
	}

	layout := "15:04"
	myStart, _ := time.Parse(layout, l.StartTime)
	myEnd, _ := time.Parse(layout, l.EndTime)
	otherStart, _ := time.Parse(layout, other.StartTime)
	otherEnd, _ := time.Parse(layout, other.EndTime)
	return myStart.Before(otherEnd) && otherStart.Before(myEnd)
}

func validateLectureTime(startTime string, endTime string) (*Lecture, error) {
	start, _ := time.Parse("15:04", startTime)
	end, _ := time.Parse("15:04", endTime)
	if start.IsZero() || end.IsZero() {
		return nil, errors.New(exception.ErrLectureTimeRequired)
	}
	if !end.After(start) {
		return nil, errors.New(exception.ErrLectureTimeOrderInvalid)
	}
	return nil, nil
}

func validateLectureDay(day Day) (*Lecture, error) {
	if day.ToKorean() == "undefined" {
		return nil, errors.New(exception.ErrLectureDayRequired)
	}
	return nil, nil
}

func validateLectureCredit(credit int) (*Lecture, error) {
	if credit < constants.LectureCreditMin || credit > constants.LectureCreditMax {
		return nil, errors.New(exception.ErrLectureCreditInvalid)
	}
	return nil, nil
}

func validateLectureCapacity(capacity int) (*Lecture, error) {
	if capacity < constants.LectureCapacityMin || capacity > constants.LectureCapacityMax {
		return nil, errors.New(exception.ErrLectureCapacityInvalid)
	}
	return nil, nil
}

func validateLectureId(id int) (*Lecture, error) {
	if id < constants.LectureIdMin || id > constants.LectureIdMax {
		return nil, errors.New(exception.ErrLectureIDInvalid)
	}
	return nil, nil
}

func validateLectureName(name string) (*Lecture, error) {
	nameLen := len([]rune(name))
	if nameLen < constants.LectureNameMin || nameLen > constants.LectureNameMax {
		return nil, errors.New(exception.ErrLectureNameRequired)
	}
	return nil, nil
}
