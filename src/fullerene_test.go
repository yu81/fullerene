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
