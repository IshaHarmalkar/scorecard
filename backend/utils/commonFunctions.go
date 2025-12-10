package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(s string) (string, error) {
	res := strings.ToLower(s)

	//replacing spaces with hyphes
	res = strings.ReplaceAll(res, " ", "-")


	//remove non-aplhanumeric characters except for hyphes
	reg, err := regexp.Compile("[^a-z0-9-]+")
	if err != nil {
		return "", err
		
	}

	res = reg.ReplaceAllLiteralString(res, "")
	return res, nil

}