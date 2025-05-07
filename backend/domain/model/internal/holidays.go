package internal

import (
	"fmt"
)

type Holidays []Day

func (r Holidays) Get(day Day) (Day, error) {
	for _, holiday := range r {
		date := holiday.Date()
		if date.Equal(day.Date()) {
			return holiday, nil
		}
	}

	return nil, fmt.Errorf("day %v does not exist", day)
}

func (r Holidays) Contains(day Day) bool {
	for _, holiday := range r {
		date := holiday.Date()
		if date.Equal(day.Date()) {
			return true
		}
	}
	return false
}

func (r Holidays) Concat(holidays Holidays) Holidays {
	temporary := Holidays{}
	temporary = append(temporary, r...)
	temporary = append(temporary, holidays...)

	return temporary
}
