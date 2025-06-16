package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract API Key from the headers HTTP REQ
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("no correct auth token in headers")
	}

	if vals[0] != "ApiKey" {
		{
			return "", errors.New("no correct first path auth token in headers")

		}
	}

	return vals[1], nil
}
