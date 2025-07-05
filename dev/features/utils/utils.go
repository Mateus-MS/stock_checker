package utils

import (
	"fmt"
	"net/http"
)

func GetQueryParam(r *http.Request, key string, required bool, defaultVal string) (string, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		if required {
			return "", fmt.Errorf("missing required query parameter: %s", key)
		}
		return defaultVal, nil
	}
	return val, nil
}
