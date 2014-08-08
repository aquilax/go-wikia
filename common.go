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

	ACTIVITY_SEGMENT                           = "Activity"
	ACTIVITY_LATEST_ACLTIVITY_SEGMENT          = "LatestActivity"
	ACTIVITY_RECENTLY_CHANGED_ARTICLES_SEGMENT = "RecentlyChangedArticles"

	ARTICLES_SEGMENT                = "Articles"
	ARTICLES_AS_SIMPLE_JSON_SEGMENT = "AsSimpleJson"
	ARTICLES_DETAILS_SEGMENT        = "Details"
	ARTICLES_LIST_SEGMENT           = "List"
	ARTICLES_MOST_LINKED_SEGMENT    = "MostLinked"
	ARTICLES_NEW_SEGMENT            = "New"
	ARTICLES_POPULAR_SEGMENT        = "Popular"
	ARTICLES_TOP_SEGMENT            = "Top"
	ARTICLES_TOP_BY_HUB_SEGMENT     = "TopByHub"

	NAVIGATION_SEGMENT      = "Navigation"
	NAVIGATION_DATA_SEGMENT = "Data"

	RELATED_PAGES_SEGMENT      = "RelatedPages"
	RELATED_PAGES_LIST_SEGMENT = "List"

	SEARCH_SEGMENT      = "Search"
	SEARCH_LIST_SEGMENT = "List"

	SEARCH_SUGGESTIONS_SEGMENT      = "SearchSuggestions"
	SEARCH_SUGGESTIONS_LIST_SEGMENT = "List"

	USER_SEGMENT         = "User"
	USER_DETAILS_SEGMENT = "Details"
)

type RequestParams map[string]string

type Revision struct {
	Id        int    `json:"id"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	Timestamp string `json:"timestamp"` //WTF: TS as string
}

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
