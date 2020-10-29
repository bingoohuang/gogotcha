package jsontime

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v := string(b)

	if len(v) >= 2 && v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]

		// 首先看是否是数字
		p, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			*t = Timestamp(time.Unix(0, p*1000000))
			return nil
		}

		v = strings.ReplaceAll(v, ",", ".")

		for _, f := range []string{
			"2006-01-02T15:04:05.000000Z",
			"2006-01-02 15:04:05.000",
		} {
			tt, err := time.Parse(f, v)
			if err == nil {
				*t = Timestamp(tt)

				return nil
			}
		}

		return errors.New("unknown timestamp format " + v)
	}

	if v == "" {
		return nil
	}

	p, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(0, p*1000000))
	return nil
}
