package utils

import (
	"net/http"
	"strconv"
)

func ToUintID(w http.ResponseWriter, r *http.Request, idv interface{}) (uint, bool) {
	var uid uint

	switch v := idv.(type) {
	case int:
		uid = uint(v)
	case int64:
		uid = uint(v)
	case uint:
		uid = v
	case uint64:
		uid = uint(v)
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			uid = uint(n)
		} else {
			return 0, false
		}
	default:
		return 0, false
	}

	return uid, true
}
