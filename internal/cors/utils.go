package cors

import "strings"

// ConvertSnakeToCamel converts snake_case to camelCase
func convertSnakeToCamel(s string) string {
	words := strings.Split(s, "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i]) // Capitalize first letter
	}
	return strings.Join(words, "")
}
