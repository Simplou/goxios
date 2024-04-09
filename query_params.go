package goxios

import "fmt"

type QueryParam struct {
	Key   string
	Value any
}

func setQueryParams(queryParams []QueryParam, url string) string {
	if len(queryParams) > 0 {
		for i, param := range queryParams {
			if i > 0 {
				url += fmt.Sprintf("&%s=%v", param.Key, param.Value)
			} else {
				url += fmt.Sprintf("?%s=%v", param.Key, param.Value)
			}
		}
	}
	return url
}
