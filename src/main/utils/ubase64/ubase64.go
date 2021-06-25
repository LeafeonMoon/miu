package ubase64

import "encoding/base64"

func BasicEncode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func BasicDecode(text string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	} else {
		return string(result), err
	}
}

func UrlEncode(text string) string {
	return base64.URLEncoding.EncodeToString([]byte(text))
}

func UrlDecode(text string) (string, error) {
	result, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	} else {
		return string(result), err
	}
}
