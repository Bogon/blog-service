package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// Iota is a special keyword in Go that represents an untyped integer constant. It is used to create a sequence of related
// constants.
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

// ToUpper "ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case."
//
// The first line of the function is the function signature. It starts with the keyword func, followed by the function
// name, the parameter list, and the return type
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower Convert a string to lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase It replaces all underscores with spaces, then capitalizes the first letter of each word, then removes all spaces
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	c := cases.Title(language.English)
	s = c.String(s) //strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase It converts a string from underscore to lower camel case.
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore If the current character is uppercase, append an underscore and the lowercase version of the character to the output.
// Otherwise, append the lowercase version of the character to the output
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
