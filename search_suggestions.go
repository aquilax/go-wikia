package gowikia

import (
	"encoding/json"
)

type SearchSuggestionsListItem struct {
	Title string `json:"title"`
}

type SearchSuggestionsListResult struct {
	Items []SearchSuggestionsListItem `json:"items"`
}

// Find suggested phrases for chosen query
// http://muppet.wikia.com/api/v1#!/SearchSuggestions/getList_get_0
func (wa *WikiaApi) SearchSuggestionsList(query string) (SearchSuggestionsListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{SEARCH_SUGGESTIONS_SEGMENT, "List"},
		RequestParams{
			"query": query,
		})
	if err != nil {
		return SearchSuggestionsListResult{}, err
	}
	var result SearchSuggestionsListResult
	return result, json.Unmarshal(jsonBlob, &result)
}
