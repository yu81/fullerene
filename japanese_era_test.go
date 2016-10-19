package fullerene

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFullerene_YearInJapaneseEra(t *testing.T) {
	d := Date(2016, 1, 1, 0, 0, 0, 0, new(time.Location))
	jy, je := d.YearInJapaneseEra()
	assert.Equal(t, jy, 28)
	assert.Equal(t, je.Name, "平成")

	d2 := Date(1800, 1, 1, 0, 0, 0, 0, new(time.Location))
	jy2, je2 := d2.YearInJapaneseEra()
	assert.Equal(t, jy2, -1)
	assert.True(t, je2 == nil)
}

func TestFullerene_DateFromJapanaseEra(t *testing.T) {
	d := DateFromJapanaseEra("平成", 28, 1, 1)
	assert.Equal(t, d.Year(), 2016)
	assert.Equal(t, int(d.Month()), 1)
	assert.Equal(t, d.Day(), 1)

	d2 := DateFromJapanaseEra("", 100, 13, 32)
	assert.Equal(t, d2, Fullerene{})
}
