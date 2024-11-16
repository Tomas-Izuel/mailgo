package notification

import (
	"fmt"
	"strings"
)

func replacePlaceholders(template string, placeholders map[string]interface{}) string {
	result := template
	for key, value := range placeholders {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}
	return result
}
