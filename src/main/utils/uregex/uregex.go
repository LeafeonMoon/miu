package uregex

import (
	"regexp"
)

func Match(regex string, val string) bool {
	match, err := regexp.MatchString(regex, val)
	return match && err == nil
}

func Compile(regex string) error {
	_, err := regexp.Compile(regex)
	return err
}
