package utils

import (
	"net/http"
	"strconv"
)

func ParsePagination(r *http.Request) (limit, offset int, err error) {
	const (
		defaultLimit = 10
		defaultPage  = 1
		maxLimit     = 100
	)

	query := r.URL.Query()

	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	page := defaultPage
	limit = defaultLimit

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			if l > 0 && l <= maxLimit {
				limit = l
			} else if l > maxLimit {
				limit = maxLimit
			}
		}
	}

	offset = (page - 1) * limit
	return limit, offset, nil
}
