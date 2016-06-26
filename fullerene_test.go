package fullerene

import (
	"github.com/stretchr/testify/assert"
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
}
