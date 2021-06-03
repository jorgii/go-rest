package restapi

import (
	"strings"
)

func ExtractBearerToken(token string) string {
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
