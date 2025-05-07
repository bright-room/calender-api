package internal

import "time"

type Reference time.Time

func (r Reference) IsBusiness() bool {
	panic("calling this method is not supported.")
}

func (r Reference) IsWeekend() bool {
	panic("calling this method is not supported.")
}

func (r Reference) IsHoliday() bool {
	panic("calling this method is not supported.")
}

func (r Reference) Summary() string {
	panic("calling this method is not supported.")
}

func (r Reference) Date() time.Time {
	return time.Time(r)
}

func (r Reference) Weekday() time.Weekday {
	return time.Time(r).Weekday()
}

func (r Reference) Description() string {
	panic("calling this method is not supported.")
}
