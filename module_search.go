package gowikia

import (
	"encoding/json"
)

type SearchType string
type SearchRank string

const (
	SEARCH_TYPE_ARTICLE SearchType = "article"
	SEARCH_TYPE_VIDEO   SearchType = "video"

	SEARCH_RANK_DEFAULT           SearchRank = "default"
	SEARCH_RANK_NEWEST            SearchRank = "newest"
	SEARCH_RANK_OLDEST            SearchRank = "oldest"
	SEARCH_RANK_RECENTLY_MODIFIED SearchRank = "recently-modified"
	SEARCH_RANK_STABLE            SearchRank = "stable"
	SEARCH_RANK_MOST_VIEWED       SearchRank = "most-viewed"
	SEARCH_RANK_FRESHEST          SearchRank = "freshest"
	SEARCH_RANK_STALEST           SearchRank = "stalest"
)

type SearchListItem struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Ns      int    `json:"ns"`
	Quality int    `json:"quality"`
}

type SearchListResult struct {
	Total        int              `json:"total"`
	Batches      int              `json:"batches"`
	CurrentBatch string           `json:"currentBatch"` // WTF String
	Next         int              `json:"next"`
	Items        []SearchListItem `json:"items"`
}

type SearchListParams struct {
	Query             string
	Type              SearchType
	Rank              SearchRank
	Limit             int
	MinArticleQuality int
	Batch             int
	Namespaces        []int
}

// Default request for Search List
func DefaultSearchListParams() SearchListParams {
	return SearchListParams{
		"test",
		SEARCH_TYPE_ARTICLE,
		SEARCH_RANK_DEFAULT,
		25,
		10,
		0,
		[]int{0},
	}
}

// Do search for given phrase
// http://muppet.wikia.com/api/v1#!/Search/getList_get_0
func (wa *WikiaApi) SearchList(params SearchListParams) (*SearchListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{SEARCH_SEGMENT, SEARCH_LIST_SEGMENT},
		RequestParams{
			"query":             params.Query,
			"type":              string(params.Type),
			"rank":              string(params.Rank),
			"limit":             intToStr(params.Limit),
			"minArticleQuality": intToStr(params.MinArticleQuality),
			"batch":             intToStr(params.Batch),
			"namespaces":        intArrToStr(params.Namespaces),
		})
	if err != nil {
		return nil, err
	}
	var result *SearchListResult
	return result, json.Unmarshal(jsonBlob, &result)
}
