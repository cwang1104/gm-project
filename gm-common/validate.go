package common

import (
	"regexp"
)

func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	matched, err := regexp.MatchString("^(?:(?:\\+|00)86)?1(?:(?:3[\\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\\d])|(?:9[1589]))\\d{8}$", mobile)
	if !matched || err != nil {
		return false
	}
	return true
}
