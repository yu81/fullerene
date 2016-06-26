package fullerene

import (
	"time"
)

type Fullerene struct {
	t time.Time
}

func Now() Fullerene {
	return Fullerene{
		t: time.Now(),
	}
}

func (fr Fullerene) Date() (year int, month time.Month, day int) {
	return fr.t.Date()
}

func (fr Fullerene) After(u Fullerene) bool {
	return fr.t.After(u.t)
}

func (fr Fullerene) Before(u Fullerene) bool {
	return fr.t.Before(u.t)
}

func (fr Fullerene) Equal(u Fullerene) bool {
	return fr.t.Equal(u.t)
}

func (fr Fullerene) IsZero() bool {
	return fr.t.IsZero()
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Fullerene {
	return Fullerene{t: time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

func (fr Fullerene) IsLeapYear() bool {
	y, _, _ := fr.Date()
	if y%4 != 0 {
		return false
	}
	if y%100 == 0 && y%400 != 0 {
		return false
	}
	return true
}

func (fr Fullerene) IsLeapDay() bool {
	_, m, d := fr.Date()
	if m == 2 && d == 29 {
		return true
	}
	return false
}

func (fr Fullerene) IsBirthday(targetTime Fullerene, beforeDayIfLeap bool) bool {
	_, m, d := fr.Date()           // birthday
	_, mm, dd := targetTime.Date() // check if it is birthday.
	if m == mm && d == dd && !fr.IsLeapDay() {
		// consider leap year.
		return true
	}

	// there are countries which a person get old at the day before leap day, and the day after in a leap year.
	return fr.isBirthdayEx(targetTime, beforeDayIfLeap)
}

func (fr Fullerene) isBirthdayEx(targetTime Fullerene, beforeDayIfLeap bool) bool {
	_, m, d := targetTime.Date()
	if targetTime.IsLeapYear() {
		return false
	}
	if beforeDayIfLeap && m == 2 && d == 28 {
		return true
	}
	if !beforeDayIfLeap && m == 3 && d == 1 {
		return true
	}
	return false
}
