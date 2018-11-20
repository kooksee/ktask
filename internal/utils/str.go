package utils

import "fmt"

func Str(a interface{}) string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%s", a)
}
