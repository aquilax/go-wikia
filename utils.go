package gowikia

import (
	"errors"
	"net/url"
	"strings"
)

const (
	API_SEGMENT = "api"
	API_VERSION = "v1"

	PATH_SEPARATOR = "/"
)

func isValidUrl(u string) (bool, error) {
	ur, err := url.Parse(u)
	if err != nil {
		return false, err
	}
	if ur.Scheme == "http" || ur.Scheme == "https" {
		return true, nil
	}
	return false, errors.New("Invalid protocol provided, only http and https are allowed")
}

func generateApiUrl(wikiaUrl, path, query string) (apiURL string) {
	parsedUrl, err := url.Parse(wikiaUrl)
	if err != nil {
		panic(err)
	}
	parsedUrl.Path = path
	parsedUrl.RawQuery = query
	return parsedUrl.String()
}

func generatePath(segments []string) string {
	segments = append([]string{API_SEGMENT, API_VERSION}, segments...)
	return strings.Join(segments, PATH_SEPARATOR)
}
