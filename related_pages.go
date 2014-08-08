package gowikia

import (
	"encoding/json"
)

type RelatedPagesListItem struct {
	Url    string `json:"url"`
	Title  string `json:"title"`
	Id     int    `json:"id"`
	ImgUrl string `json:"imgUrl"` // WTF: CamelCase
	Text   string `json:"text"`
}

type RelatedPagesListResult struct {
	Items    []RelatedPagesListItem `json:"items"`
	Basepath string                 `json:"basepath"`
}

// Get pages related to a given article ID
// http://muppet.wikia.com/api/v1#!/RelatedPages/getList_get_0
func (wa *WikiaApi) RelatedPagesList(ids []int, limit int) (RelatedPagesListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{RELATED_PAGES_SEGMENT, "List"},
		RequestParams{
			"ids":   intArrToStr(ids),
			"limit": intToStr(limit),
		})
	if err != nil {
		return RelatedPagesListResult{}, err
	}
	var result RelatedPagesListResult
	return result, json.Unmarshal(jsonBlob, &result)
}
