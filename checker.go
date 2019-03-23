package jsontagchecker

import "strings"

func IsJsonTag(tag string) bool {
	return strings.Contains(tag, "json")
}

func Checker(fieldName string, tag string) bool {
	split := strings.Split(tag, "\"")
	if len(split) == 3 {
		tagName := split[1]
		tagSplit := strings.Split(tagName, "_")
		join := strings.Join(tagSplit, "")
		return strings.ToLower(fieldName) == join
	}
	return false
}
