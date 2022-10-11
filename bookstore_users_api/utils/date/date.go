package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	// time format, standard time zone
	return time.Now().UTC()
}

func GetNowString() string {
	// Z is time zone
	// dd-mm-yyyy
	// now.Format("02-01-2006T15:04:05Z")
	// mm-dd-yyyy
	// now.Format("01-02-2006T15:04:05Z")
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
