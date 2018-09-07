package nemtime

import (
	"testing"
)

var (
	incorrectInputFormat = "J03L5ux"
	startDates           = map[string]string{
		"2018-01-01": "2018-01-01T00:00:00+10:00",
		"2018-02-01": "2018-02-01T00:00:00+10:00",
		"2018-03-01": "2018-03-01T00:00:00+10:00",
		"2018-04-01": "2018-04-01T00:00:00+10:00",
		"2018-05-01": "2018-05-01T00:00:00+10:00",
		"2018-06-01": "2018-06-01T00:00:00+10:00",
	}
	endDates = map[string]string{
		"2018-06-30": "2018-07-01T00:00:00+10:00",
		"2018-07-31": "2018-08-01T00:00:00+10:00",
		"2018-08-31": "2018-09-01T00:00:00+10:00",
		"2018-09-30": "2018-10-01T00:00:00+10:00",
		"2018-10-31": "2018-11-01T00:00:00+10:00",
		"2018-11-30": "2018-12-01T00:00:00+10:00",
		"2018-12-31": "2019-01-01T00:00:00+10:00",
	}
)

func TestStartDateStringToNEMTimeString(t *testing.T) {
	for k, v := range startDates {
		out, err := FromStartDateString(k)
		if err != nil {
			t.Error(err)
		}

		if out.String() != v {
			t.Errorf("Expected output to be %v, got %v", v, out)
		}
	}
}

func TestEndDateStringToNEMTimeString(t *testing.T) {
	for k, v := range endDates {
		out, err := FromEndDateString(k)
		if err != nil {
			t.Error(err)
		}

		if out.String() != v {
			t.Errorf("Expected output to be %v, got %v", v, out)
		}
	}
}

func TestErrors(t *testing.T) {
	_, err := FromStartDateString(incorrectInputFormat)
	if err == nil {
		t.Error("Expected an error, got none!")
	}

	_, err = FromEndDateString(incorrectInputFormat)
	if err == nil {
		t.Error("Expected an error, got none!")
	}
}
