package golazone

import (
	"fmt"
	"strings"

	godiacritics "gopkg.in/Regis24GmbH/go-diacritics.v2"
)

// GetZone returns a timezone string from a city name and country code.
// If no match is found, the function returns "UTC".
func GetZone(cityName, countryCode string) string {
	cityName = strings.ToLower(strings.TrimSpace(godiacritics.Normalize((cityName))))
	cityName = strings.Join(strings.Fields(cityName), " ")
	countryCode = strings.TrimSpace(countryCode)

	out := ""
	if cityName != "" && countryCode != "" {
		out = cityTimezoneMap[fmt.Sprintf("%s_%s", cityName, strings.ToUpper(countryCode))]
		if out != "" {
			return out
		}
	}

	if out == "" && cityName != "" {
		try1City := strings.ReplaceAll(cityName, " ", "")
		out = cityAsciiToTimezone[try1City]
	}

	if out == "" && countryCode != "" {
		out = iso2ToTimezone[strings.ToLower(countryCode)]
	}

	if out == "" {
		out = "UTC"
	}

	return out
}
