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
		return out, throw
	}

	isValidNumber := phonenumbers.IsValidNumber(data)
	if !isValidNumber {
		throw = fmt.Errorf("phone number is not valid: %s", phone)

		return out, throw
	}

	out = data

	return out, throw
}

func FormatPhone(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return out, throw
	}

	out = phonenumbers.Format(phoneNumber, phonenumbers.E164)

	return out, throw
}

func FormatPhoneToCountryCode(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return out, throw
	}

	out = strconv.Itoa(int(phoneNumber.GetCountryCode()))

	return out, throw
}

func FormatPhoneToSignificantNumber(phone string) (out string, throw error) {
	phoneNumber, throw := Parse(phone)
	if nil != throw {
		return out, throw
	}

	out = phonenumbers.GetNationalSignificantNumber(phoneNumber)

	return out, throw
}
