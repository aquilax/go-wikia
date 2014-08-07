package gowikia

import (
	"net/url"
	"strings"
)

const (
	API_SEGMENT = "api"
	API_VERSION = "v1"

	PATH_SEPARATOR = "/"
)

func generateApiUrl(wikiaUrl, path, query string) (apiURL string, err error) {
	parsedUrl, err := url.Parse(wikiaUrl)
	if err != nil {
		return "", err
	}
	parsedUrl.Path = path
	parsedUrl.RawQuery = query
	return parsedUrl.String(), nil
}

func generatePath(segments []string) string {
	segments = append([]string{API_SEGMENT, API_VERSION}, segments...)
	return strings.Join(segments, PATH_SEPARATOR)
}
