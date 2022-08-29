package books

import (
	"fmt"
	"strings"
	"time"
)

type ReleaseDate struct {
	time.Time
}

const releaseDateLayout = "2006-01-02"

func NewReleaseDate(date string) (*ReleaseDate, error) {
	layout := "2006-01-02"

	datetime, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}
	return &ReleaseDate{Time: datetime}, nil
}

func (d *ReleaseDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse(releaseDateLayout, s)

	return
}

func (d *ReleaseDate) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%q", d.Time.Format(releaseDateLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (d *ReleaseDate) IsSet() bool {
	return d.UnixNano() != nilTime
}
