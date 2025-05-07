package internal

import (
	"time"
)

type NationalHoliday struct {
	date        time.Time
	description string
}

func NewNationalHoliday(date time.Time, description string) NationalHoliday {
	return NationalHoliday{
		date:        date,
		description: description,
	}
}

func (r NationalHoliday) IsBusiness() bool {
	return false
}

func (r NationalHoliday) IsWeekend() bool {
	return false
}

func (r NationalHoliday) IsHoliday() bool {
	return true
}

func (r NationalHoliday) Summary() string {
	return "祝日"
}

func (r NationalHoliday) Date() time.Time {
	return r.date
}

func (r NationalHoliday) Weekday() time.Weekday {
	return r.date.Weekday()
}

func (r NationalHoliday) Description() string {
	return r.description
}
