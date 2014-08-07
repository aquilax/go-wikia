package gowikia

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

const (
	API_SEGMENT = "api"
	API_VERSION = "v1"

	SEPARATOR_PATH = "/"
	SEPARATOR_INT  = ","
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
	return strings.Join(segments, SEPARATOR_PATH)
}

func generateQuery(hash map[string]string) string {
	values := url.Values{}
	for key, val := range hash {
		values.Set(key, val)
	}
	return values.Encode()
}

func intArrToStr(numbers []int) string {
	var strArray []string
	for _, number := range numbers {
		strArray = append(strArray, strconv.Itoa(number))
	}
	return strings.Join(strArray, SEPARATOR_INT)
}

func intToStr(number int) string {
	return strconv.Itoa(number)
}
