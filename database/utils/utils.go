package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ObjectKeys(obj interface{}) []string {
	keys := make([]string, len(obj.(map[string]string)))
	i := 0
	for k := range obj.(map[string]string) {
		keys[i] = k
		i++
	}
	return keys
}

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func ToProperCase(s string) string {
	return cases.Title(language.Und).String(s)
}