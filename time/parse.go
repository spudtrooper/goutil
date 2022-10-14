package time

import (
	"time"

	"github.com/pkg/errors"
)

var (
	layouts = []string{
		"2006-01-02",
		"2006/01/02",
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
	}
)

// Parse is a drop-in replacement for time.Parse that tries a bunch of layouts
// before giving up.
func Parse(value string) (time.Time, error) {
	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.Errorf("could not parse time from %s", value)
}

// ParseInLocation is a drop-in replacement for time.ParseInLocation that tries a
// bunch of layouts before giving up.
func ParseInLocation(value string, loc *time.Location) (time.Time, error) {
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.Errorf("could not parse time from %s", value)
}

// ParseDuration is a drop-in replacement for time.ParseDuration that tries a
// bunch of layouts before giving up.
func ParseDuration(value string) (time.Duration, error) {
	d, err := time.ParseDuration(value)
	if err == nil {
		return d, nil
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t.Sub(time.Time{}), nil
		}
	}
	return 0, errors.Errorf("could not parse duration from %s", value)
}

// ParseDurationInLocation is a drop-in replacement for
// time.ParseDurationInLocation that tries a bunch of layouts before giving up.
func ParseDurationInLocation(value string, loc *time.Location) (time.Duration, error) {
	d, err := time.ParseDuration(value)
	if err == nil {
		return d, nil
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t.Sub(time.Time{}), nil
		}
	}
	return 0, errors.Errorf("could not parse duration from %s", value)
}
