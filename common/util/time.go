package util

import "time"

// GetJSTTimeStr get time yyyymmdd str.
func GetJSTTimeStr(dayThrethold int) string {
	// get past 7 days alerts
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().UTC().In(jst).AddDate(0, 0, -dayThrethold).Format("20060102")
	return nowJST
}

// GetNowJSTTimeStr get current time yyyy/mm/dd hh:mm:ss
func GetNowJSTTimeStr() string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().UTC().In(jst).Format("2006/01/02 15:04:05")
	return nowJST
}
