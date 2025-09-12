package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Valid phone number
	phoneNumber, err := Parse("+12125550100")
	assert.NoError(t, err)
	assert.NotNil(t, phoneNumber)
	assert.Equal(t, int32(1), phoneNumber.GetCountryCode())
	assert.Equal(t, uint64(2125550100), phoneNumber.GetNationalNumber())

	// Invalid phone number (too short, but parsable by libphonenumber, invalid by IsValidNumber)
	phoneNumber, err = Parse("+1555")
	assert.Error(t, err)
	assert.Nil(t, phoneNumber)
	assert.EqualError(t, err, "phone number is not valid: +1555")

	// Invalid phone number (non-numeric, not parsable by libphonenumber)
	phoneNumber, err = Parse("abc")
	assert.Error(t, err)
	assert.Nil(t, phoneNumber)
	assert.Contains(t, err.Error(), "the phone number supplied is not a number")

	// Empty string (not parsable by libphonenumber)
	phoneNumber, err = Parse("")
	assert.Error(t, err)
	assert.Nil(t, phoneNumber)
	assert.Contains(t, err.Error(), "the phone number supplied is not a number")
}

func TestFormatPhone(t *testing.T) {
	// Valid phone number
	formattedPhone, err := FormatPhone("+12125550100")
	assert.NoError(t, err)
	assert.Equal(t, "+12125550100", formattedPhone)

	// Invalid phone number
	formattedPhone, err = FormatPhone("123")
	assert.Error(t, err)
	assert.Empty(t, formattedPhone)
	assert.Contains(t, err.Error(), "invalid country code")
}

func TestFormatPhoneToCountryCode(t *testing.T) {
	// Valid phone number
	countryCode, err := FormatPhoneToCountryCode("+12125550100")
	assert.NoError(t, err)
	assert.Equal(t, "1", countryCode)

	// Invalid phone number
	countryCode, err = FormatPhoneToCountryCode("123")
	assert.Error(t, err)
	assert.Empty(t, countryCode)
	assert.Contains(t, err.Error(), "invalid country code")
}

func TestFormatPhoneToSignificantNumber(t *testing.T) {
	// Valid phone number
	significantNumber, err := FormatPhoneToSignificantNumber("+12125550100")
	assert.NoError(t, err)
	assert.Equal(t, "2125550100", significantNumber)

	// Invalid phone number
	significantNumber, err = FormatPhoneToSignificantNumber("123")
	assert.Error(t, err)
	assert.Empty(t, significantNumber)
	assert.Contains(t, err.Error(), "invalid country code")
}
