package Examples

import (
	"errors"
	"fmt"
	"time"
)

// Note: Go time layout uses the reference time Mon Jan 2 15:04:05 MST 2006.
const (
	D_YYMMDD            string = "06-01-02"                // yy-MM-dd
	D_DDMMyy            string = "02-01-06"                // dd-MM-yy
	D_YYMMDD_N          string = "06-Jan-02"               // yy-MMM-dd
	D_DDMMyy_N          string = "02-Jan-06"               // dd-MMM-yy
	D_YYMMDDHHMMA_N     string = "06-Jan-02, 03:04PM"      // yy-MMM-dd, hh:mma
	D_DDMMyyHHMMA_N     string = "02-Jan-06, 03:04PM"      // dd-MMM-yy, hh:mma
	S_YYMMDD            string = "06/01/02"                // yy/MM/dd
	S_DDMMyy            string = "02/01/06"                // dd/MM/yy
	S_YYMMDDHHMMA       string = "06/01/02, 03:04PM"       // yy/MM/dd, hh:mma
	S_DDMMyyHHMMA       string = "02/01/06, 03:04PM"       // dd/MM/yy, hh:mma
	S_YYMMDDHHMMA_N     string = "06/Jan/02, 03:04PM"      // yy/MMM/dd, hh:mma
	S_DDMMyyHHMMA_N     string = "02/Jan/06, 03:04PM"      // dd/MMM/yy, hh:mma
	D_YYYYMMDD          string = "2006-01-02"              // yyyy-MM-dd
	D_DDMMYYYY          string = "02-01-2006"              // dd-MM-yyyy
	D_YYYYMMDDHHMMA     string = "2006-01-02, 03:04PM"     // yyyy-MM-dd, hh:mma
	D_DDMMYYYYHHMMA     string = "02-01-2006, 03:04PM"     // dd-MM-yyyy, hh:mma
	D_YYYYMMDD_N        string = "2006-Jan-02"             // yyyy-MMM-dd
	D_DDMMYYYY_N        string = "02-Jan-2006"             // dd-MMM-yyyy
	D_YYYYMMDDHHMMA_N   string = "2006-Jan-02, 03:04PM"    // yyyy-MMM-dd, hh:mma
	D_DDMMYYYYHHMMA_N   string = "02-Jan-2006, 03:04PM"    // dd-MMM-yyyy, hh:mma
	S_YYYYMMDD          string = "2006/01/02"              // yyyy/MM/dd
	S_DDMMYYYY          string = "02/01/2006"              // dd/MM/yyyy
	S_YYYYMMDDHHMMA     string = "2006/01/02, 03:04PM"     // yyyy/MM/dd, hh:mma
	S_DDMMYYYYHHMMA     string = "02/01/2006, 03:04PM"     // dd/MM/yyyy, hh:mma
	S_YYYYMMDDHHMMA_N   string = "2006/Jan/02, 03:04PM"    // yyyy/MMM/dd, hh:mma
	S_DDMMYYYYHHMMA_N   string = "02/Jan/2006, 03:04PM"    // dd/MMM/yyyy, hh:mma
	D_YYMMDDHHMMSSA_N   string = "06-Jan-02, 03:04:05PM"   // yy-MMM-dd, hh:mm:ssa
	D_DDMMyyHHMMSSA_N   string = "02-Jan-06, 03:04:05PM"   // dd-MMM-yy, hh:mm:ssa
	S_YYMMDDHHMMSSA     string = "06/01/02, 03:04:05PM"    // yy/MM/dd, hh:mm:ssa
	S_DDMMyyHHMMSSA     string = "02/01/06, 03:04:05PM"    // dd/MM/yy, hh:mm:ssa
	S_YYMMDDHHMMSSA_N   string = "06/Jan/02, 03:04:05PM"   // yy/MMM/dd, hh:mm:ssa
	S_DDMMyyHHMMSSA_N   string = "02/Jan/06, 03:04:05PM"   // dd/MMM/yy, hh:mm:ssa
	D_YYYYMMDDHHMMSSA   string = "2006-01-02, 03:04:05PM"  // yyyy-MM-dd, hh:mm:ssa
	D_DDMMYYYYHHMMSSA   string = "02-01-2006, 03:04:05PM"  // dd-MM-yyyy, hh:mm:ssa
	D_YYYYMMDDHHMMSSA_N string = "2006-Jan-02, 03:04:05PM" // yyyy-MMM-dd, hh:mm:ssa
	D_DDMMYYYYHHMMSSA_N string = "02-Jan-2006, 03:04:05PM" // dd-MMM-yyyy, hh:mm:ssa
	S_YYYYMMDDHHMMSSA   string = "2006/01/02, 03:04:05PM"  // yyyy/MM/dd, hh:mm:ssa
	S_DDMMYYYYHHMMSSA   string = "02/01/2006, 03:04:05PM"  // dd/MM/yyyy, hh:mm:ssa
	S_YYYYMMDDHHMMSSA_N string = "2006/Jan/02, 03:04:05PM" // yyyy/MMM/dd, hh:mm:ssa
	S_DDMMYYYYHHMMSSA_N string = "02/Jan/2006, 03:04:05PM" // dd/MMM/yyyy, hh:mm:ssa
	HHMMA               string = "03:04PM"                 // hh:mma
	HHMM                string = "03:04"                   // hh:mm
	HHMMSSA             string = "03:04:05PM"              // hh:mm:ssa
	HHMMSS              string = "03:04:05"                // hh:mm:ss
)

// DateHelper encapsulates core date and time manipulation functions.
type DateHelper struct{}

// NewDateHelper creates and returns a new DateHelper instance.
func NewDateHelper() *DateHelper {
	return &DateHelper{}
}

// PrettifyDate formats a timestamp (in milliseconds) into a human-readable string.
// If the timestamp is today, it returns time in "03:04 PM" format, otherwise in "02 Jan 03:04 PM" format.
func (dh *DateHelper) PrettifyDate(timeMillis int64) string {
	t := time.UnixMilli(timeMillis)
	var layout string
	if isToday(t) {
		layout = HHMMA
	} else {
		layout = "02 Jan 03:04 PM"
	}
	return t.Format(layout)
}

// PrettifyDateFromString converts a timestamp string (milliseconds) and formats it.
func (dh *DateHelper) PrettifyDateFromString(timestampStr string) (string, error) {
	timestamp, err := parseInt64(timestampStr)
	if err != nil {
		return "", err
	}
	return dh.PrettifyDate(timestamp), nil
}

// GetDateOnly parses a date string in "02/01/2006" (dd/MM/yyyy) format and returns the timestamp in milliseconds.
func (dh *DateHelper) GetDateOnly(dateStr string) (int64, error) {
	layout := S_DDMMYYYY
	t, err := time.ParseInLocation(layout, dateStr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

// GetDateOnlyFromTime formats a timestamp (in milliseconds) into "02/01/2006" (dd/MM/yyyy) format.
func (dh *DateHelper) GetDateOnlyFromTime(timeMillis int64) string {
	layout := S_DDMMYYYY
	t := time.UnixMilli(timeMillis)
	return t.Format(layout)
}

// GetDateAndTime formats a timestamp (in milliseconds) into "02/01/2006, 03:04 PM" format.
func (dh *DateHelper) GetDateAndTime(timeMillis int64) string {
	layout := S_DDMMYYYYHHMMA
	t := time.UnixMilli(timeMillis)
	return t.Format(layout)
}

// GetTimeOnly formats a timestamp (in milliseconds) into "03:04 PM" (hh:mm a) format.
func (dh *DateHelper) GetTimeOnly(timeMillis int64) string {
	layout := HHMMA
	t := time.UnixMilli(timeMillis)
	return t.Format(layout)
}

// GetTodayWithTime returns the current date and time in "02/01/2006 15:04:05" format.
func (dh *DateHelper) GetTodayWithTime() string {
	layout := "02/01/2006 15:04:05"
	return time.Now().Format(layout)
}

// GetToday returns today's date in "02/01/2006" format.
func (dh *DateHelper) GetToday() string {
	layout := S_DDMMYYYY
	return time.Now().Format(layout)
}

// GetTomorrow returns tomorrow's date in "02/01/2006" format.
func (dh *DateHelper) GetTomorrow() string {
	today, _ := time.Parse(S_DDMMYYYY, dh.GetToday())
	tomorrow := today.AddDate(0, 0, 1)
	return tomorrow.Format(S_DDMMYYYY)
}

// getDurationBetween parses two date strings with the provided layout and returns the duration between them.
func (dh *DateHelper) getDurationBetween(oldStr, newStr string, layout string) (time.Duration, error) {
	t1, err := time.ParseInLocation(layout, oldStr, time.Local)
	if err != nil {
		return 0, err
	}
	t2, err := time.ParseInLocation(layout, newStr, time.Local)
	if err != nil {
		return 0, err
	}
	return t1.Sub(t2), nil
}

// GetDaysBetweenTwoDate returns the number of days between two date strings based on the given DateFormat.
func (dh *DateHelper) GetDaysBetweenTwoDate(oldStr, newStr string, df string) (int64, error) {
	dur, err := dh.getDurationBetween(oldStr, newStr, df)
	if err != nil {
		return 0, err
	}
	return int64(dur.Hours() / 24), nil
}

// GetHoursBetweenTwoDate returns the number of hours between two date strings based on the given DateFormat.
func (dh *DateHelper) GetHoursBetweenTwoDate(oldStr, newStr string, df string) (int64, error) {
	dur, err := dh.getDurationBetween(oldStr, newStr, df)
	if err != nil {
		return 0, err
	}
	return int64(dur.Hours()), nil
}

// GetMinutesBetweenTwoDates returns the number of minutes between two date strings based on the given DateFormat.
func (dh *DateHelper) GetMinutesBetweenTwoDates(oldStr, newStr string, df string) (int64, error) {
	dur, err := dh.getDurationBetween(oldStr, newStr, df)
	if err != nil {
		return 0, err
	}
	return int64(dur.Minutes()), nil
}

// ParseAnyDate attempts to parse the input date string using all defined DateFormats.
// Returns the timestamp in milliseconds for the first successful parse.
func (dh *DateHelper) ParseAnyDate(dateStr string) (int64, error) {
	formats := []string{
		D_YYMMDD, D_DDMMyy, D_YYMMDD_N, D_DDMMyy_N,
		D_YYMMDDHHMMA_N, D_DDMMyyHHMMA_N, S_YYMMDD, S_DDMMyy,
		S_YYMMDDHHMMA, S_DDMMyyHHMMA, S_YYMMDDHHMMA_N, S_DDMMyyHHMMA_N,
		D_YYYYMMDD, D_DDMMYYYY, D_YYYYMMDDHHMMA, D_DDMMYYYYHHMMA,
		D_YYYYMMDD_N, D_DDMMYYYY_N, D_YYYYMMDDHHMMA_N, D_DDMMYYYYHHMMA_N,
		S_YYYYMMDD, S_DDMMYYYY, S_YYYYMMDDHHMMA, S_DDMMYYYYHHMMA,
		S_YYYYMMDDHHMMA_N, S_DDMMYYYYHHMMA_N, D_YYMMDDHHMMSSA_N, D_DDMMyyHHMMSSA_N,
		S_YYMMDDHHMMSSA, S_DDMMyyHHMMSSA, S_YYMMDDHHMMSSA_N, S_DDMMyyHHMMSSA_N,
		D_YYYYMMDDHHMMSSA, D_DDMMYYYYHHMMSSA, D_YYYYMMDDHHMMSSA_N, D_DDMMYYYYHHMMSSA_N,
		S_YYYYMMDDHHMMSSA, S_DDMMYYYYHHMMSSA, S_YYYYMMDDHHMMSSA_N, S_DDMMYYYYHHMMSSA_N,
		HHMMA, HHMM, HHMMSSA, HHMMSS,
	}

	for _, format := range formats {
		t, err := time.ParseInLocation(format, dateStr, time.Local)
		if err == nil {
			return t.UnixMilli(), nil
		}
	}
	return 0, errors.New("unable to parse date with any known format")
}

// ParseDate parses the input date string using the specified DateFormat and returns the timestamp in milliseconds.
func (dh *DateHelper) ParseDate(dateStr string, df string) (int64, error) {
	t, err := time.ParseInLocation(df, dateStr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.UnixMilli(), nil
}

// GetDesiredFormat returns the current date formatted with the specified DateFormat.
func (dh *DateHelper) GetDesiredFormat(df string) string {
	return time.Now().Format(df)
}

// GetDesiredFormatFromTime returns the formatted date for the given timestamp (in milliseconds) using the specified DateFormat.
func (dh *DateHelper) GetDesiredFormatFromTime(df string, timeMillis int64) string {
	t := time.UnixMilli(timeMillis)
	return t.Format(df)
}

// isToday checks if the provided time is on the same day as the current time.
func isToday(t time.Time) bool {
	now := time.Now()
	y1, m1, d1 := now.Date()
	y2, m2, d2 := t.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// parseInt64 converts a string into an int64 value.
func parseInt64(s string) (int64, error) {
	var i int64
	_, err := fmt.Sscan(s, &i)
	if err != nil {
		return 0, err
	}
	return i, nil
}
