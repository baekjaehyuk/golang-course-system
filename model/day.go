package model

import "golang-course-registration/common/constants"

type Day string

const (
	Monday    Day = "MON"
	Tuesday   Day = "TUE"
	Wednesday Day = "WED"
	Thursday  Day = "THU"
	Friday    Day = "FRI"
)

func (d Day) ToKorean() string {
	switch d {
	case Monday:
		return constants.MON
	case Tuesday:
		return constants.TUE
	case Wednesday:
		return constants.WED
	case Thursday:
		return constants.THU
	case Friday:
		return constants.FRI
	default:
		return constants.Undefined
	}
}
