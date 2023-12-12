package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type TimeStamp time.Time

const TimeFormat = "2006-01-02 15:04:05"

func (t TimeStamp) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (ts *TimeStamp) ToTime() time.Time {
	return time.Time(*ts)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	// 空值不进行解析
	if len(data) == 2 {
		*ts = TimeStamp(time.Time{})
		return errors.New("时间为nil")
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*ts = TimeStamp(now)
	return err
}

func (ts TimeStamp) ToString() string {
	return ts.ToTime().Format("2006-01-02 15:04:05")
}

func (ts TimeStamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(ts)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (ts *TimeStamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*ts = TimeStamp(value)
		return nil
	}
	//i, err = strconv.ParseInt(sc, 10, 64)

	return fmt.Errorf("can not convert %v to timestamp", v)
}
