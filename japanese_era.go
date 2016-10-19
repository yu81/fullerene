package fullerene

import "time"

type JapaneseEra struct {
	Name         string
	KanaName     string
	HiraganaName string
	Start        Fullerene
	End          Fullerene
}

var JapaneseEraList = []JapaneseEra{
	{
		Name:         "平成",
		KanaName:     "ヘイセイ",
		HiraganaName: "へいせい",
		Start:        Date(1989, 1, 8, 0, 0, 0, 0, new(time.Location)),
		End:          Date(9999, 12, 31, 0, 0, 0, 0, new(time.Location)),
	},
	{
		Name:         "昭和",
		KanaName:     "ショウワ",
		HiraganaName: "しょうわ",
		Start:        Date(1926, 12, 25, 0, 0, 0, 0, new(time.Location)),
		End:          Date(1989, 1, 8, 0, 0, 0, 0, new(time.Location)),
	},
	{
		Name:         "大正",
		KanaName:     "タイショウ",
		HiraganaName: "たいしょう",
		Start:        Date(1912, 7, 30, 0, 0, 0, 0, new(time.Location)),
		End:          Date(1926, 12, 25, 0, 0, 0, 0, new(time.Location)),
	},
	{
		Name:         "明治",
		KanaName:     "メイジ",
		HiraganaName: "めいじ",
		Start:        Date(1868, 1, 1, 0, 0, 0, 0, new(time.Location)),
		End:          Date(1912, 7, 30, 0, 0, 0, 0, new(time.Location)),
	},
}

// YearInJapaneseEra returns year in japanese era (ex. 2016 -> 28 )
func (fr Fullerene) YearInJapaneseEra() (int, *JapaneseEra) {
	for _, jEra := range JapaneseEraList {
		if fr.After(jEra.Start) && fr.Before(jEra.End) {
			return fr.Year() - jEra.Start.Year() + 1, &jEra
		}
	}
	return -1, nil
}

func DateFromJapanaseEra(eraName string, japaneseYear int, month time.Month, day int) Fullerene {
	for _, jEra := range JapaneseEraList {
		if jEra.Name == eraName || jEra.HiraganaName == eraName || jEra.KanaName == eraName {
			return Date(jEra.Start.Year()-1+japaneseYear, month, day, 0, 0, 0, 0, new(time.Location))
		}
	}
	return Fullerene{}
}
