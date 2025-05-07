package internal

import "time"

type ClosedDay struct {
	date        time.Time
	description string
}

func NewClosedDay(date time.Time, description string) ClosedDay {
	return ClosedDay{
		date:        date,
		description: description,
	}
}

func (r ClosedDay) IsBusiness() bool {
	return false
}

func (r ClosedDay) IsWeekend() bool {
	return false
}

func (r ClosedDay) IsHoliday() bool {
	return true
}

func (r ClosedDay) Summary() string {
	return "休業日"
}

func (r ClosedDay) Date() time.Time {
	return r.date
}

func (r ClosedDay) Weekday() time.Weekday {
	return r.date.Weekday()
}

func (r ClosedDay) Description() string {
	return r.description
}
