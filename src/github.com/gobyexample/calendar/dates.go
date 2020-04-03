package calendar

import "errors"

// Date making day, month, year, not exportable
type Date struct {
	day   int
	month int
	year  int
}

// SetYear blah
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid date")
	}
	d.year = year
	return nil
}

// Year Getter
func (d *Date) Year() int {
	return d.year
}

// SetMonth blah
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// Month Getter
func (d *Date) Month() int {
	return d.month
}

// SetDay blah
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// Day Getter
func (d *Date) Day() int {
	return d.day
}
