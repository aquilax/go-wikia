package gowikia

import (
	"encoding/json"
)

type UserDetailsItem struct {
	UserId        int    `json:"user_id"`
	Title         string `json:"title"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	NumberOfEdits int    `json:"numberofedits"` // WTF: lowercase
	Avatar        string `json:"avatar"`
}

type UserDetailsResult struct {
	Items    []UserDetailsItem `json:"items"`
	Basepath string            `json:"basepath"`
}

// Find suggested phrases for chosen query
// http://muppet.wikia.com/api/v1#!/SearchSuggestions/getList_get_0
func (wa *WikiaApi) UserDetails(ids []int, size int) (UserDetailsResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{USER_SEGMENT, "Details"},
		RequestParams{
			"ids":  intArrToStr(ids),
			"size": intToStr(size),
		})
	if err != nil {
		return UserDetailsResult{}, err
	}
	var result UserDetailsResult
	return result, json.Unmarshal(jsonBlob, &result)
}
