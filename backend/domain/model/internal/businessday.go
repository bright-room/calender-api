package internal

import "time"

type Businessday struct {
	date time.Time
}

func (r Businessday) IsBusiness() bool {
	return true
}

func (r Businessday) IsWeekend() bool {
	return false
}

func (r Businessday) IsHoliday() bool {
	return false
}

func (r Businessday) Summary() string {
	return "営業日"
}

func (r Businessday) Date() time.Time {
	return r.date
}

func (r Businessday) Weekday() time.Weekday {
	return r.date.Weekday()
}

func (r Businessday) Description() string {
	return ""
}
