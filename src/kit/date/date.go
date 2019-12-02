package date

import (
	"bytes"
	"errors"
	"time"
)

// Date is a formatted time which contemplates days, months and years. Provides methods to parse
// from and to other formats and has Before, After... kind of operations.
// Unmarshals methods will work with "2006-01-02" formatted times.
//
// Zero value Date is Jesus birthday: 0001-01-01
type Date struct {
	v string
	t time.Time
}

const dateLayoutDashYYYYMMDD = "2006-01-02"

// DateNow builds a Date from today
func DateNow() Date {
	t := time.Now().UTC()
	return Date{
		v: t.Format("2006-01-02"),
		t: t,
	}
}

// DateFromTime builds a Date from a time.Time
func DateFromTime(t time.Time) Date {
	var b [len(dateLayoutDashYYYYMMDD)]byte
	buf := b[:0]
	buf = t.AppendFormat(buf, dateLayoutDashYYYYMMDD)
	return Date{
		v: string(buf),
		t: t,
	}
}

// Time returns the 'time.Time' value corresponding to a Date. Note that hours of 'Time' will always be 0 since
// a Date only arranges from year to day
func (d Date) Time() (time.Time, error) {
	if d.v == "" && d.t.IsZero() {
		return d.toTime()
	}
	return d.t, nil
}

// FormatToDashYYYYMMDD formats the Date to 2006-01-02 format
func (d Date) FormatToDashYYYYMMDD() string {
	if d.v == "" {
		return "0001-01-01"
	}
	return d.v
}

// DateFromDashYYYYMMDDFormat builds a Date from a 2006-01-02 formatted time with no format validation.
// For format validations use the safe version (which is 100 times slower)
func DateFromDashYYYYMMDDFormat(s string) (Date, error) {
	if len(s) != 10 {
		return Date{}, errors.New("date '" + s + "' must be in the '2006-01-02' format")
	}
	return Date{v: s}, nil
}

// DateFromDashYYYYMMDDFormatSafe builds a Date from a 2006-01-02 formatted time validating the input
// This is 100 times slower than the unsafe version
func DateFromDashYYYYMMDDFormatSafe(s string) (Date, error) {
	// time.Parse to validate 's'
	t, err := time.Parse(dateLayoutDashYYYYMMDD, s)
	if err != nil {
		return Date{}, err
	}
	d := Date{
		v: s,
		t: t,
	}
	return d, nil
}

// FormatToShort formats the Date to 060102 format
func (d Date) FormatToShort() string {
	if d.v == "" {
		return "010101"
	}
	return string(d.v[2:4] + d.v[5:7] + d.v[8:10])
}

// DateFromShortFormat builds a Date from a 060102 formatted time
func DateFromShortFormat(s string) (Date, error) {
	if len(s) != 6 {
		return Date{}, errors.New("date is not in '060102' format: " + s)
	}
	return Date{v: "20" + s[0:2] + "-" + s[2:4] + "-" + s[4:6]}, nil
}

// FormatToSlashDDMMYYYY formats a Date to 02/01/2006 format
func (d Date) FormatToSlashDDMMYYYY() string {
	if d.v == "" {
		return "01/01/0001"
	}
	return string(d.v[8:10] + "/" + d.v[5:7] + "/" + d.v[0:4])
}

// DateFromSlashDDMMYYYYFormat builds a Date from a 02/01/2006 formatted time
func DateFromSlashDDMMYYYYFormat(s string) (Date, error) {
	if len(s) != 10 {
		return Date{}, errors.New("date is not in '02/01/2006' format: " + s)
	}
	return Date{v: s[6:10] + "-" + s[3:5] + "-" + s[0:2]}, nil
}

// Equal evaluates if this date equals to u
func (d Date) Equal(u Date) bool {
	return d.v == u.v
}

// Before evaluates if this date is a date before u
func (d Date) Before(u Date) bool {
	return d.v < u.v
}

// BeforeOrEqual evaluates if this date is before or equals to u
func (d Date) BeforeOrEqual(u Date) bool {
	return d.v <= u.v
}

// After evaluates if this date is a date after u
func (d Date) After(u Date) bool {
	return d.v > u.v
}

// AfterOrEqual evaluates if this date is after or equals to u
func (d Date) AfterOrEqual(u Date) bool {
	return d.v >= u.v
}

// DatesDaysCompare returns an in
func DatesDaysCompare(high, low *Date) (int, error) {
	err := datesSetTimes(high, low)
	if err != nil {
		return 0, err
	}
	dur := high.t.Sub(low.t)
	days := int(dur.Hours() / 24)
	return days, nil
}

func datesSetTimes(high, low *Date) error {
	err := dateNilCheck(high, low)
	if err != nil {
		return err
	}
	if high.t.IsZero() {
		_, err = high.toTime()
		if err != nil {
			return err
		}
	}
	if low.t.IsZero() {
		_, err = low.toTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func dateNilCheck(high, low *Date) error {
	if high == nil {
		return errors.New("high not initialized")
	}
	if low == nil {
		return errors.New("low not initialized")
	}
	return nil
}

func (d *Date) toTime() (time.Time, error) {
	t, err := time.Parse(dateLayoutDashYYYYMMDD, d.v)
	if err != nil {
		return time.Time{}, err
	}
	d.t = t
	return t, nil
}

// MarshalJSON supports json.Marshaler interface. Marshals to a "2006-01-02" formatted string
func (d Date) MarshalJSON() ([]byte, error) {
	// append quotes
	return []byte("\"" + d.FormatToDashYYYYMMDD() + "\""), nil
}

// UnmarshalJSON supports json.Unmarshal interface. Expects a "2006-01-02" formatted string
func (d *Date) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return errors.New("unmarshal error, null is not a valid Date")
	}
	if len(data) <= 2 {
		return errors.New("unmarshal error, date is empty")
	}
	// remove quotes
	s := string(data[1 : len(data)-1])
	tmp, err := DateFromDashYYYYMMDDFormatSafe(s)
	if err != nil {
		return errors.New("unmarshal error, " + err.Error())
	}
	*d = tmp
	return nil
}
