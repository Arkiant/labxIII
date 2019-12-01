package date

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestDate_DashYYYYMMDDFormat(t *testing.T) {
	_, err := DateFromDashYYYYMMDDFormat("44")
	assert.Error(t, err)

	date := "2019-12-10"
	d1, err := DateFromDashYYYYMMDDFormat(date)
	assert.NoError(t, err)
	assert.Equal(t, date, d1.FormatToDashYYYYMMDD())
}

func TestDate_DashYYYYMMDDFormatSafe(t *testing.T) {
	d, err := DateFromDashYYYYMMDDFormatSafe("1992-11-23")
	assert.NoError(t, err)
	assert.Equal(t, "1992-11-23", d.FormatToDashYYYYMMDD())

	_, err = DateFromDashYYYYMMDDFormatSafe("23/11/1992")
	assert.Error(t, err)

	_, err = DateFromDashYYYYMMDDFormatSafe("")
	assert.Error(t, err)

	_, err = DateFromDashYYYYMMDDFormatSafe("1992-1-23")
	assert.Error(t, err)

	_, err = DateFromDashYYYYMMDDFormatSafe("1992-13-23")
	assert.Error(t, err)
}

func TestDate_ShortFormat(t *testing.T) {
	_, err := DateFromShortFormat("44")
	assert.Error(t, err)

	date := "190101"
	d1, err := DateFromShortFormat(date)
	assert.NoError(t, err)
	assert.Equal(t, date, d1.FormatToShort())
}

func TestDate_SlashDDMMYYYYFormat(t *testing.T) {
	_, err := DateFromSlashDDMMYYYYFormat("44")
	assert.Error(t, err)

	date := "23/11/1992"
	d1, err := DateFromSlashDDMMYYYYFormat(date)
	assert.NoError(t, err)
	assert.Equal(t, date, d1.FormatToSlashDDMMYYYY())
}

func TestDate_EqualsDashAndShortFormat(t *testing.T) {
	date, err := DateFromDashYYYYMMDDFormat("2019-01-01")
	assert.NoError(t, err)
	actual := date.FormatToShort()
	exp := "190101"
	assert.Equal(t, string(exp), actual)
}

func TestDate_EqualsSlashAndShortFormat(t *testing.T) {
	date, err := DateFromDashYYYYMMDDFormat("2019-02-01")
	assert.NoError(t, err)
	actual := date.FormatToSlashDDMMYYYY()
	exp := "01/02/2019"
	assert.Equal(t, exp, actual)
}

func TestDate_Equal(t *testing.T) {
	start := "2019-07-03"
	end := "2019-07-03"
	sDate, err := DateFromDashYYYYMMDDFormat(start)
	assert.NoError(t, err)
	eDate, err := DateFromDashYYYYMMDDFormat(end)
	assert.NoError(t, err)
	assert.True(t, sDate.Equal(eDate))
	assert.True(t, sDate.AfterOrEqual(eDate))
	assert.True(t, sDate.BeforeOrEqual(eDate))
}

func TestDate_Before_Day(t *testing.T) {
	start := "2019-07-02"
	end := "2019-07-03"
	testDateBefore(t, start, end)
}

func TestDate_Before_Month(t *testing.T) {
	start := "2019-05-23"
	end := "2019-12-23"
	testDateBefore(t, start, end)

	start = "2019-05-24"
	end = "2019-12-23"
	testDateBefore(t, start, end)

	start = "2019-05-23"
	end = "2019-12-24"
	testDateBefore(t, start, end)

}

func TestDate_Before_Year(t *testing.T) {
	start := "2016-07-03"
	end := "2019-07-03"
	testDateBefore(t, start, end)

	start = "2016-06-03"
	end = "2019-07-03"
	testDateBefore(t, start, end)

	start = "2016-07-03"
	end = "2019-08-03"
	testDateBefore(t, start, end)
}

func testDateBefore(t *testing.T, start, end string) {
	sDate, err := DateFromDashYYYYMMDDFormat(start)
	assert.NoError(t, err)
	eDate, err := DateFromDashYYYYMMDDFormat(end)
	assert.NoError(t, err)
	assert.True(t, sDate.Before(eDate))
	assert.True(t, sDate.BeforeOrEqual(eDate))
}

func TestDate_Before_Day_False(t *testing.T) {
	start := "2019-05-24"
	end := "2019-05-23"
	testDateNotBefore(t, start, end)
}

func TestDate_Before_Month_Error(t *testing.T) {
	start := "2019-12-23"
	end := "2019-05-24"
	testDateNotBefore(t, start, end)

	start = "2019-12-23"
	end = "2019-05-23"
	testDateNotBefore(t, start, end)

	start = "2019-12-24"
	end = "2019-05-23"
	testDateNotBefore(t, start, end)
}

func TestDate_Before_Year_Error(t *testing.T) {
	start := "2020-12-23"
	end := "2019-05-24"
	testDateNotBefore(t, start, end)

	start = "2020-05-23"
	end = "2019-05-23"
	testDateNotBefore(t, start, end)

	start = "2020-04-23"
	end = "2019-05-23"
	testDateNotBefore(t, start, end)
}

func testDateNotBefore(t *testing.T, start, end string) {
	sDate, err := DateFromDashYYYYMMDDFormat(start)
	assert.NoError(t, err)
	eDate, err := DateFromDashYYYYMMDDFormat(end)
	assert.NoError(t, err)
	assert.False(t, sDate.Before(eDate))
	assert.False(t, sDate.BeforeOrEqual(eDate))
}

func TestDate_After_Day(t *testing.T) {
	start := "2019-07-04"
	end := "2019-07-03"
	testDateAfter(t, start, end)
}

func TestDate_After_Month(t *testing.T) {
	start := "2019-08-03"
	end := "2019-07-03"
	testDateAfter(t, start, end)
}

func TestDate_After_Year(t *testing.T) {
	start := "2020-07-03"
	end := "2019-07-03"
	testDateAfter(t, start, end)
}

func testDateAfter(t *testing.T, start, end string) {
	sDate, err := DateFromDashYYYYMMDDFormat(start)
	assert.NoError(t, err)
	eDate, err := DateFromDashYYYYMMDDFormat(end)
	assert.NoError(t, err)
	assert.True(t, sDate.After(eDate), start, end, "start: %s, end: %s", start, end)
	assert.True(t, sDate.AfterOrEqual(eDate), "start: %s, end: %s", start, end)
}

func TestDate_DatesDaysCompare(t *testing.T) {
	start := "2019-07-03"
	end := "2019-07-04"
	sd, err := DateFromDashYYYYMMDDFormat(start)
	assert.NoError(t, err)
	ed, err := DateFromDashYYYYMMDDFormat(end)
	assert.NoError(t, err)

	days, err := DatesDaysCompare(&ed, &sd)
	assert.NoError(t, err)
	assert.Equal(t, 1, days)

	days, err = DatesDaysCompare(&sd, &ed)
	assert.NoError(t, err)
	assert.Equal(t, -1, days)

	_, err = DatesDaysCompare(nil, nil)
	assert.Error(t, err)

	_, err = DatesDaysCompare(&sd, nil)
	assert.Error(t, err)

	_, err = DatesDaysCompare(nil, &sd)
	assert.Error(t, err)
}

func TestDate_MarshalUnmarshal(t *testing.T) {
	tn := []byte(`"2019-05-22"`)
	var d Date
	err := json.Unmarshal(tn, &d)
	assert.NoError(t, err)

	b, err := json.Marshal(d)
	assert.NoError(t, err)

	success, err := gomega.MatchJSON(tn).Match(b)
	assert.NoError(t, err)
	assert.True(t, success)
}

func TestDate_Unmarshal_Error(t *testing.T) {
	tn := []byte(`""`)
	var d Date
	err := json.Unmarshal(tn, &d)
	assert.Error(t, err)

	tn = []byte(`null`)
	err = json.Unmarshal(tn, &d)
	assert.Error(t, err)

	tn = []byte(`"23/11/1992"`)
	err = json.Unmarshal(tn, &d)
	assert.Error(t, err)

	tn = []byte(`"2019-5-28"`)
	err = json.Unmarshal(tn, &d)
	assert.Error(t, err)
}

func BenchmarkDate_SafeUnsafe(b *testing.B) {
	tn := "2019-07-03"
	b.Run("Safe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DateFromDashYYYYMMDDFormatSafe(tn)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("Unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := DateFromDashYYYYMMDDFormat(tn)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkDate_BeforeAfter(b *testing.B) {
	tn1, _ := DateFromDashYYYYMMDDFormat("2019-05-24")
	tn2, _ := DateFromDashYYYYMMDDFormat("2019-05-25")
	b.Run("before", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tn1.Before(tn2)
		}
	})
	b.Run("After", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tn1.After(tn2)
		}
	})
}

// Bench the consume when create instance from string of date
func BenchmarkDate_Create(b *testing.B) {
	checkinDate := "2019-07-03"
	b.Run("time.Parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r, err := time.Parse("2006-01-02", checkinDate)
			if err != nil {
				b.Fatal(err)
			}
			_ = r
		}
	})
	b.Run("DateFromDashYYYYMMDDFormat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r, err := DateFromDashYYYYMMDDFormat(checkinDate)
			if err != nil {
				b.Fatal(err)
			}
			_ = r
		}
	})
}

func BenchmarkDate_Operations(b *testing.B) {
	checkinDate := "2019-01-02"
	from := "2019-01-01"
	to := "2019-01-03"
	b.Run("time.Time", func(b *testing.B) {
		date, err := time.Parse("2006-01-02", checkinDate)
		if err != nil {
			b.Fatal(err)
		}
		dateTo, err := time.Parse("2006-01-02", from)
		if err != nil {
			b.Fatal(err)
		}
		dateFrom, err := time.Parse("2006-01-02", to)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r := timeCheckDate(date, &dateFrom, &dateTo)
			_ = r
		}
	})

	b.Run("Date", func(b *testing.B) {
		date, err := DateFromDashYYYYMMDDFormat(checkinDate)
		if err != nil {
			b.Fatal(err)
		}
		dateFrom, err := DateFromDashYYYYMMDDFormat(from)
		if err != nil {
			b.Fatal(err)
		}
		dateTo, err := DateFromDashYYYYMMDDFormat(to)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r := dateStr2CheckDate(date, &dateFrom, &dateTo)
			_ = r
		}
	})
}

func timeCheckDate(dateCandidate time.Time, from *time.Time, to *time.Time) bool {
	if from == nil && to == nil {
		return true
	}
	if from == nil && to != nil {
		return beforeOrEqual(dateCandidate, *to)
	}
	if to == nil && from != nil {
		return afterOrEqual(dateCandidate, *from)
	}
	if dateCandidate.Equal(*from) || dateCandidate.Equal(*to) {
		return true
	}
	return beforeOrEqual(dateCandidate, *to) && afterOrEqual(dateCandidate, *from)
}

func afterOrEqual(t time.Time, u time.Time) bool {
	return t.After(u) || t.Equal(u)
}

func beforeOrEqual(t time.Time, u time.Time) bool {
	return t.Before(u) || t.Equal(u)
}

func dateStr2CheckDate(dateCandidate Date, from *Date, to *Date) bool {
	if from == nil && to == nil {
		return true
	}
	if from == nil && to != nil {
		return dateCandidate.BeforeOrEqual(*to)
	}
	if to == nil && from != nil {
		return dateCandidate.AfterOrEqual(*from)
	}
	if dateCandidate.Equal(*from) || dateCandidate.Equal(*to) {
		return true
	}
	return dateCandidate.BeforeOrEqual(*to) && dateCandidate.AfterOrEqual(*from)
}

func BenchmarkDate_DatesDaysCompare(b *testing.B) {
	start := "2019-07-03"
	end := "2019-07-04"

	b.Run("from initialized time", func(b *testing.B) {
		sd, err := DateFromDashYYYYMMDDFormatSafe(start)
		assert.NoError(b, err)
		ed, err := DateFromDashYYYYMMDDFormatSafe(end)
		assert.NoError(b, err)
		for i := 0; i < b.N; i++ {
			sdtmp := sd
			edtmp := ed
			days, err := DatesDaysCompare(&sdtmp, &edtmp)
			if err != nil {
				b.Fatal(err)
			}
			_ = days
		}
	})
	b.Run("initializing time", func(b *testing.B) {
		sd, err := DateFromDashYYYYMMDDFormat(start)
		assert.NoError(b, err)
		ed, err := DateFromDashYYYYMMDDFormat(end)
		assert.NoError(b, err)
		for i := 0; i < b.N; i++ {
			sdtmp := sd
			edtmp := ed
			days, err := DatesDaysCompare(&sdtmp, &edtmp)
			if err != nil {
				b.Fatal(err)
			}
			_ = days
		}
	})
}
