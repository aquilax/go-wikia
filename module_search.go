package gowikia

import (
	"encoding/json"
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
	CurrentBatch int              `json:"currentBatch"`
	Next         int              `json:"next"`
	Items        []SearchListItem `json:"items"`
}

type SearchListParams struct {
	Query             string
	Type              string
	Rank              string
	Limit             int
	MinArticleQuality int
	Batch             int
	Namespaces        []int
}

// Do search for given phrase
// http://muppet.wikia.com/api/v1#!/Search/getList_get_0
func (wa *WikiaApi) SearchList(params SearchListParams) (*SearchListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{SEARCH_SEGMENT, "List"},
		RequestParams{
			"query":             params.Query,
			"type":              params.Type,
			"rank":              params.Rank,
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
