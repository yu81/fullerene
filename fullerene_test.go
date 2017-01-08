package fullerene

import (
	"fmt"
	libage "github.com/bearbin/go-age"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	mt, tt := Now(), time.Now()
	mty, mtm, mtd := mt.Date()
	tty, ttm, ttd := tt.Date()
	assert.Equal(t, mty, tty)
	assert.Equal(t, mtm, ttm)
	assert.Equal(t, mtd, ttd)

	assert.True(t, mt.Equal(mt))
}

func TestFullerene_After(t *testing.T) {
	mt, tt := Now(), time.Now()
	assert.False(t, mt.After(Fullerene{t: tt}))
}

func TestFullerene_Before(t *testing.T) {
	mt, tt := Now(), time.Now()
	assert.True(t, mt.Before(Fullerene{t: tt}))
}

func TestFullerene_IsZero(t *testing.T) {
	assert.True(t, Fullerene{}.IsZero())
	assert.False(t, Now().IsZero())
}

func TestFullerene_Equal(t *testing.T) {
	mt := Now()
	assert.True(t, mt.Equal(mt))
}

func TestFullerene_IsLeapYear(t *testing.T) {
	fr := Date(2016, 1, 1, 0, 0, 0, 0, &time.Location{})
	assert.True(t, fr.IsLeapYear())
	fr2 := Date(2000, 1, 1, 0, 0, 0, 0, &time.Location{})
	assert.True(t, fr2.IsLeapYear())
	fr3 := Date(1999, 1, 1, 0, 0, 0, 0, &time.Location{})
	assert.False(t, fr3.IsLeapYear())
	fr4 := Date(2100, 1, 1, 0, 0, 0, 0, &time.Location{})
	assert.False(t, fr4.IsLeapYear())
}

func TestFullerene_IsLeapDay(t *testing.T) {
	fr := Date(2016, 2, 29, 0, 0, 0, 0, &time.Location{})
	assert.True(t, fr.IsLeapDay())

	fr2 := Date(2016, 2, 28, 0, 0, 0, 0, &time.Location{})
	assert.False(t, fr2.IsLeapDay())
}

func TestFullerene_IsBirthday(t *testing.T) {
	emptyLocation := time.Location{}
	fr := Date(2014, 11, 18, 0, 0, 0, 0, &emptyLocation)
	assert.True(t, fr.IsBirthday(Date(1981, 11, 18, 0, 0, 0, 0, &emptyLocation), false))
	assert.False(t, fr.IsBirthday(Date(1981, 11, 19, 0, 0, 0, 0, &emptyLocation), false))

	frLeap := Date(1980, 2, 29, 0, 0, 0, 0, &emptyLocation)
	assert.False(t, frLeap.IsBirthday(Date(2015, 2, 28, 0, 0, 0, 0, &emptyLocation), false))
	assert.True(t, frLeap.IsBirthday(Date(2015, 3, 1, 0, 0, 0, 0, &emptyLocation), false))
	assert.True(t, frLeap.IsBirthday(Date(2015, 2, 28, 0, 0, 0, 0, &emptyLocation), true))
	assert.False(t, frLeap.IsBirthday(Date(2015, 3, 1, 0, 0, 0, 0, &emptyLocation), true))
	assert.False(t, frLeap.IsBirthday(Date(2016, 3, 1, 0, 0, 0, 0, &emptyLocation), true))
	assert.False(t, frLeap.IsBirthday(Date(2016, 3, 1, 0, 0, 0, 0, &emptyLocation), false))
	tm := time.Now()
	tm.Weekday()
	fr.isBirthdayEx(frLeap, false)
}

func TestFullerene_Age(t *testing.T) {
	type Case struct {
		Birthday    Fullerene
		ExpectedAge int
	}
	targetDate := Date(2016, 10, 2, 0, 0, 0, 0, &time.Location{})
	tests := []Case{
		{Date(2015, 1, 1, 0, 0, 0, 0, &time.Location{}), 1},
		{Date(2014, 12, 31, 0, 0, 0, 0, &time.Location{}), 1},
		{Date(1999, 7, 16, 0, 0, 0, 0, &time.Location{}), 17},
		{Date(1988, 2, 29, 0, 0, 0, 0, &time.Location{}), 28},
		{Date(1988, 3, 3, 0, 0, 0, 0, &time.Location{}), 28},
		{Date(1988, 10, 1, 0, 0, 0, 0, &time.Location{}), 27},
		{Date(1988, 10, 2, 0, 0, 0, 0, &time.Location{}), 28},
		{Date(1988, 10, 3, 0, 0, 0, 0, &time.Location{}), 28},
	}

	for _, d := range tests {
		assert.Equal(t, d.ExpectedAge, d.Birthday.Age(targetDate), d.Birthday.String())
	}
}

func BenchmarkFullerene_Date(b *testing.B) {
	location := &time.Location{}
	for i := 0; i < b.N; i++ {
		Date(2016, 1, 1, 0, 0, 0, 0, location)
	}
}

func BenchmarkFullerene_Date2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Date(2016, 1, 1, 0, 0, 0, 0, &time.Location{})
	}
}

func BenchmarkFullerene_Date3(b *testing.B) {
	b.StopTimer()
	fr := Date(2016, 1, 1, 0, 0, 0, 0, &time.Location{})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = fr.Year()
	}
}

func BenchmarkFullerene_Date4(b *testing.B) {
	b.StopTimer()
	fr := Date(2016, 1, 1, 0, 0, 0, 0, &time.Location{})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fr.Month(), fr.Day()
	}
}

func BenchmarkFullerene_Date5(b *testing.B) {
	b.StopTimer()
	fr := Date(2016, 1, 1, 0, 0, 0, 0, &time.Location{})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = fr.Date()
	}
}

func Benchmark_Date(b *testing.B) {
	b.StopTimer()
	location := &time.Location{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		time.Date(2016, 1, 1, 0, 0, 0, 0, location)
	}
}

func Benchmark_Fullerene_Age(b *testing.B) {
	b.StopTimer()
	f := Date(2015, 1, 1, 0, 0, 0, 0, &time.Location{})
	b.StartTimer()
	now := Now()
	for i := 0; i < b.N; i++ {
		f.Age(now)
	}
}

func AgeFromString(birthday, targetDay string) int {
	y1, _ := strconv.Atoi(birthday[0:4])
	y2, _ := strconv.Atoi(targetDay[0:4])
	age := y2 - y1
	m1, _ := strconv.Atoi(birthday[4:6])
	m2, _ := strconv.Atoi(targetDay[4:6])
	if m1 < m2 {
		return age
	}
	if m1 > m2 {
		return age - 1
	}
	d1, _ := strconv.Atoi(birthday[6:8])
	d2, _ := strconv.Atoi(targetDay[6:8])
	if d1 < d2 {
		return age - 1
	}
	return age
}

func TestSimpleAge(t *testing.T) {
	assert.Equal(t, 6, AgeFromString("20100101", "20161020"))
	assert.Equal(t, 6, AgeFromString("20100101", "20160101"))
	assert.Equal(t, 5, AgeFromString("20100101", "20151231"))
}

func BenchmarkSimpleAge(b *testing.B) {
	b.StopTimer()
	f1 := Date(2010, 10, 21, 0, 0, 0, 0, &time.Location{})
	f2 := Date(2010, 10, 22, 0, 0, 0, 0, &time.Location{})
	d1 := (f1.String())[0:12]
	d2 := (f2.String())[0:12]
	fmt.Println(d1)
	fmt.Println(d2)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		AgeFromString(d1, d2)
	}
}

func BenchmarkFullerene_GoAge(b *testing.B) {
	b.StopTimer()
	t := time.Now().AddDate(0,0,-1)
	b.StartTimer()
	for i:=0;i<b.N;i++ {
		libage.Age(t)
	}
}
