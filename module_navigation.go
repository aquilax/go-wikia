package gowikia

import (
	"encoding/json"
)

type NavigationDataItem struct {
	Text     string               `json:"text"`
	Href     string               `json:"href"`
	Children []NavigationDataItem `json:"children"`
}

type NavigationDataList struct {
	Wikia []NavigationDataItem `json:"wikia"`
	Wiki  []NavigationDataItem `json:"wiki"`
}

type NavigationDataResult struct {
	Navigation NavigationDataList `json:"navigation"`
}

// Get wiki navigation links (the main menu of given wiki)
// http://muppet.wikia.com/api/v1#!/Navigation/getData_get_0
func (wa *WikiaApi) NavigationData() (*NavigationDataResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{NAVIGATION_SEGMENT, NAVIGATION_DATA_SEGMENT},
		RequestParams{})
	if err != nil {
		return nil, err
	}
	var result *NavigationDataResult
	return result, json.Unmarshal(jsonBlob, &result)
}
