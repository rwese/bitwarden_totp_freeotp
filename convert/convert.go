package convert

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Convert takes an input JSON data and returns the converted data
func Convert(r io.Reader) ([]byte, error) {
	var jsonData JSONData
	err := json.NewDecoder(r).Decode(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	items := jsonData.Items

	var otpURLs []string
	for _, item := range items {
		if item.Login.TOTP != "" {
			otpURL := fmt.Sprintf("otpauth://totp/%s?secret=%s", item.Name, item.Login.TOTP)
			otpURLs = append(otpURLs, otpURL)
		}
	}

	return []byte(strings.Join(otpURLs, "\n")), nil
}
