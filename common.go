package gowikia

import (
	"io/ioutil"
	"net/http"
)

const (
	API_SEGMENT = "api"
	API_VERSION = "v1"

	SEPARATOR_PATH  = "/"
	SEPARATOR_ARRAY = ","

	ACTIVITY_SEGMENT           = "Activity"
	ARTICLES_SEGMENT           = "Articles"
	NAVIGATION_SEGMENT         = "Navigation"
	RELATED_PAGES_SEGMENT      = "RelatedPages"
	SEARCH_SEGMENT             = "Search"
	SEARCH_SUGGESTIONS_SEGMENT = "SearchSuggestions"
	USER_SEGMENT               = "User"
)

type RequestParams map[string]string

type OriginalDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func getJsonBlob(wikiaUrl string, segments []string, params RequestParams) ([]byte, error) {
	path := generatePath(segments)
	query := generateQuery(params)
	url := generateApiUrl(wikiaUrl, path, query)
	return fetchUrl(url)
}

func fetchUrl(u string) ([]byte, error) {
	response, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
