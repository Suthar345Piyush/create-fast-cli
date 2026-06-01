// some helper functions - all are string based will used by writer and renderer while scaffolding the project

package utils

import (
	"strings"
	"unicode"
)

// slugfunc - function will convert the app name to lowercase letters with hyphens in between them

// My CLI app -> my-cli-app

func SlugFunc(projectName string) string {
	projectName = strings.ToLower(strings.TrimSpace(projectName))

	var sb strings.Builder

	hyphen := false

	for _, r := range projectName {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
			hyphen = false
		} else if !hyphen && sb.Len() > 0 {
			sb.WriteString("-")
			hyphen = true
		}
	}

	result := sb.String()

	return strings.TrimRight(result, "-")

}

// pascal case function to convert the slug or snake-case string to pascal case
// my-cli-app -> MyCliApp

func PascalCase(projectName string) string {

	words := strings.FieldsFunc(projectName, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	var sb strings.Builder

	for _, w := range words {
		if len(w) == 0 {
			continue
		}

		runes := []rune(w)

		sb.WriteRune(unicode.ToUpper(runes[0]))

		for _, r := range runes[:1] {
			sb.WriteRune(unicode.ToLower(r))
		}
	}

	return sb.String()

}

// snake-case or slug to camel case
// my-cli-app -> muCliApp

func CamelCase(projectName string) string {

	p := PascalCase(projectName)

	if len(p) == 0 {
		return p
	}

	runes := []rune(p)

	return string(unicode.ToLower(runes[0])) + string(runes[1:])

}

// removing file extension - main.go.tmpl -> main.go

func RemoveExt(filename string) string {

	idx := strings.LastIndex(filename, ".")

	if idx < 0 {
		return filename
	}

	return filename[:idx]
}

// if string is empty or contains only whitespaces

func IsEmpty(st string) bool {
	return strings.TrimSpace(st) == ""
}
