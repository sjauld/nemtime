package nemtime

import (
	"time"
)

// According the Australian Energy Market Operator, everything happens in
// UTC+10:00
var (
	NEMTZOffset     = 10 * time.Hour
	NEMTZCorrection = time.Duration(-NEMTZOffset)
	NEMTimeZone     = time.FixedZone("NEM", int(NEMTZOffset.Seconds()))

	// We'll accept a reasonably standard date format as input
	DateLayout = "2006-01-02"
)

// NEMTime represents a point in time as understood by a simple string, from the
// perspective of the Australian Energy Market Operator
type NEMTime struct {
	time.Time
}

// FromStartDateString takes a simple date string as an input and returns a
// pointer to a NEMTime representing the start of that day
func FromStartDateString(s string) (*NEMTime, error) {
	baseTime, err := time.Parse(DateLayout, s)
	if err != nil {
		return nil, err
	}

	return &NEMTime{baseTime.In(NEMTimeZone).Add(NEMTZCorrection)}, nil
}

// FromEndDateString takes a simple date string as an input and returns a
// pointer to a NEMTime representing the end of that day
func FromEndDateString(s string) (*NEMTime, error) {
	baseTime, err := time.Parse(DateLayout, s)
	if err != nil {
		return nil, err
	}

	return &NEMTime{baseTime.In(NEMTimeZone).Add(NEMTZCorrection).Add(time.Duration(24 * time.Hour))}, nil
}

// String represents a NEMTime in an RFC3339 string
func (n *NEMTime) String() string {
	return n.Format(time.RFC3339)
}
