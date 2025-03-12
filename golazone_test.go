package golazone

import (
	"testing"
)

func TestGetZone(t *testing.T) {
	tests := []struct {
		name        string
		cityName    string
		countryCode string
		expected    string
	}{
		// ----- Direct Matches -----
		{
			name:        "DirectMatch: Đà Nẵng + VN",
			cityName:    "Đà Nẵng",
			countryCode: "VN",
			expected:    "Asia/Ho_Chi_Minh",
		},
		{
			name:        "DirectMatch: New York + US",
			cityName:    "New York",
			countryCode: "US",
			expected:    "America/New_York",
		},
		{
			name:        "DirectMatch: Paris + FR",
			cityName:    "Paris",
			countryCode: "FR",
			expected:    "Europe/Paris",
		},

		// ----- Input Normalization -----
		{
			name:        "Normalize: mixed case + extra spaces",
			cityName:    "  ĐÀ   NẵNG  ",
			countryCode: " vn ",
			expected:    "Asia/Ho_Chi_Minh",
		},
		{
			name:        "Normalize: multiple spaces + lower case",
			cityName:    "los   angeles",
			countryCode: "Us",
			expected:    "America/Los_Angeles",
		},

		// ----- Fallback: cityAsciiToTimezone -----
		{
			name:        "Fallback: cityAsciiToTimezone no country",
			cityName:    "Đà Nẵng",
			countryCode: "",
			expected:    "Asia/Ho_Chi_Minh",
		},
		{
			name:        "Fallback: cityAsciiToTimezone multiple spaces",
			cityName:    "New     York",
			countryCode: "",
			expected:    "America/New_York",
		},

		// ----- Fallback: iso2ToTimezone -----
		{
			name:        "Fallback: iso2ToTimezone country only",
			cityName:    "",
			countryCode: "fr",
			expected:    "Europe/Paris",
		},
		{
			name:        "Fallback: iso2ToTimezone country upper case",
			cityName:    "",
			countryCode: "VN",
			expected:    "Asia/Ho_Chi_Minh",
		},

		// ----- Defaults to UTC -----
		{
			name:        "Default: unknown city and country",
			cityName:    "UnknownCity",
			countryCode: "ZZ",
			expected:    "UTC",
		},
		{
			name:        "Default: empty inputs",
			cityName:    "",
			countryCode: "",
			expected:    "UTC",
		},

		// ----- Edge Cases -----
		{
			name:        "Edge: numeric city name",
			cityName:    "12345",
			countryCode: "US",
			expected:    "America/New_York", // fallback on country
		},
		{
			name:        "Edge: special characters in city name",
			cityName:    "!@#$%",
			countryCode: "VN",
			expected:    "Asia/Ho_Chi_Minh", // fallback on country
		},
		{
			name:        "Edge: known city with incorrect country code",
			cityName:    "Đà Nẵng",
			countryCode: "ZZ",
			expected:    "Asia/Ho_Chi_Minh", // fallback cityAsciiToTimezone
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetZone(tt.cityName, tt.countryCode)
			if result != tt.expected {
				t.Errorf("GetZone(%q, %q) = %q; want %q",
					tt.cityName, tt.countryCode, result, tt.expected)
			}
		})
	}
}
