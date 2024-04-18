package goxios

import (
	"fmt"
	netUrl "net/url"
)

// QueryParam represents a single key-value pair for query parameters.
type QueryParam struct {
	Key   string // The key of the query parameter.
	Value any    // The value of the query parameter.
}

// setQueryParams constructs a URL string with the given query parameters.
// It takes a slice of QueryParam and a base URL string as input.
// It returns the URL string with query parameters appended.
func setQueryParams(queryParams []QueryParam, url string) string {
	// Check if there are any query parameters to append.
	if len(queryParams) > 0 {
		// Iterate over each query parameter.
		for i, param := range queryParams {
			value := fmt.Sprintf("%v", param.Value)
			// If it's not the first parameter, append '&' before the key-value pair.
			if i > 0 {
				url += fmt.Sprintf("&%s=%s", param.Key, netUrl.QueryEscape(value))
			} else {
				// If it's the first parameter, append '?' before the key-value pair.
				url += fmt.Sprintf("?%s=%s", param.Key, netUrl.QueryEscape(value))
			}
		}
	}
	// Return the constructed URL string with query parameters.
	return url
}
