package baseutils

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// ExtractNullInt
// extract a null.Int out of a interface that might not be int64
func ExtractNullInt(data any) (v null.Int) {
	switch data := data.(type) {
	case float64:
		v = null.NewInt(int64(data), true)
	}

	return
}

// ExtractNullTime
// extract a null.Time out of a string interface by parsing it in various formats
func ExtractNullTime(data any) (v null.Time) {
	switch data := data.(type) {
	case string:
		date, err := time.Parse("2006-01-02T15:04:05", data)
		if err != nil {
			date, err = time.Parse("2006-01-02T15:04:05Z", data)
			if err != nil {
				date, err = time.Parse("2006-01-02T15:04:05-03:00", data)
				if err != nil {
					date, err = time.Parse("2006-01-02T15:04:05+03:00", data)
					if err != nil {
						return
					}
				}
			}
		}

		v = null.NewTime(date, true)
	}

	return v
}

// ExtractNullString
// extract a null.String out of a interface that might not be string
func ExtractNullString(data any) (v null.String) {
	switch data := data.(type) {
	case string:
		if data != "" {
			return null.NewString(data, true)
		}
	}

	return
}

// ExtractStringArray
// extract a array of strings out of an interface
func ExtractStringArray(data any) (v []string) {
	switch data.(type) {
	case []any:
		for _, e := range data.([]any) {
			switch e := e.(type) {
			case string:
				v = append(v, e)
			}
		}
	}

	return
}
