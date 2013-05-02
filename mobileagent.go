package mobileagent

import (
	"regexp"
)

var reDoCoMoRE = regexp.MustCompile(`^DoCoMo/\d\.\d[ /]`)
var reJPhoneRE = regexp.MustCompile(`^(?i:J-PHONE/\d\.\d)`)
var reVodafoneRE = regexp.MustCompile(`^Vodafone/\d\.\d`)
var reVodafoneMotRE = regexp.MustCompile(`^MOT-`)
var reSoftBankRE = regexp.MustCompile(`^SoftBank/\d\.\d`)
var reSoftBankCrawlerRE = regexp.MustCompile(`^Nokia[^/]+/\d\.\d`)
var reEZwebRE = regexp.MustCompile(`^(?:KDDI-[A-Z]+\d+[A-Z]? )?UP\.Browser/`)
var reAirHRE = regexp.MustCompile(`^Mozilla/3\.0\((?:WILLCOM|DDIPOCKET);`)

func IsDoCoMo(userAgent string) bool {
	return reDoCoMoRE.MatchString(userAgent)
}

func IsJPhone(userAgent string) bool {
	return reJPhoneRE.MatchString(userAgent)
}

func IsVodaphone(userAgent string) bool {
	return reVodafoneRE.MatchString(userAgent) ||
		reVodafoneMotRE.MatchString(userAgent)
}

func IsThirdforce(userAgent string) bool {
	return reJPhoneRE.MatchString(userAgent) ||
		reVodafoneRE.MatchString(userAgent) ||
		reVodafoneMotRE.MatchString(userAgent) ||
		reSoftBankRE.MatchString(userAgent) ||
		reSoftBankCrawlerRE.MatchString(userAgent)
}

func IsSoftBank(userAgent string) bool {
	return IsThirdforce(userAgent)
}

func IsEZWeb(userAgent string) bool {
	return reEZwebRE.MatchString(userAgent)
}

func IsAirH(userAgent string) bool {
	return reAirHRE.MatchString(userAgent)
}

func IsMobile(userAgent string) bool {
	return IsDoCoMo(userAgent) ||
		IsJPhone(userAgent) ||
		IsVodaphone(userAgent) ||
		IsThirdforce(userAgent) ||
		IsEZWeb(userAgent) ||
		IsAirH(userAgent)
}
