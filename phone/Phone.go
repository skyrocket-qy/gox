package phone

import (
	"fmt"
	"strconv"

	"github.com/nyaruka/phonenumbers"
)

func Parse(phone string) (out *phonenumbers.PhoneNumber, throw error) {
	region := ""
	data, throw := phonenumbers.Parse(phone, region)
	if nil != throw {
		return
	}

	isValidNumber := phonenumbers.IsValidNumber(data)
	if !isValidNumber {
		throw = fmt.Errorf("phone number is not valid: %s", phone)
		return
	}
	out = data
	return
}

func FormatPhone(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return
	}
	out = phonenumbers.Format(phoneNumber, phonenumbers.E164)
	return
}

func FormatPhoneToCountryCode(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return
	}

	out = strconv.Itoa(int(phoneNumber.GetCountryCode()))
	return
}

func FormatPhoneToSignificantNumber(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return
	}

	out = phonenumbers.GetNationalSignificantNumber(phoneNumber)
	return
}
