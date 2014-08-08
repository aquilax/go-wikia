package gowikia

import (
	"encoding/json"
	"strconv"
)

type ActivityResultItem struct {
	Article    int `json:"article"`
	User       int `json:"user"`
	RevisionId int `json:"revisionId"`
	Timestamp  int `json:"timestamp"`
}

type ActivityResult struct {
	Items    []ActivityResultItem `json:"items"`
	Basepath string               `json:"basepath"`
}

type ActivityRequest struct {
	Limit           int
	Namespaces      []int
	AllowDuplicates bool
}

// Provides default values for LatestActivity function
func DefaultActivityRequest() ActivityRequest {
	return ActivityRequest{10, []int{0}, true}
}

// Get latest activity information
// http://muppet.wikia.com/api/v1#!/Activity/getLatestActivity_get_0
func (wa *WikiaApi) ActivityLatestActivity(params ActivityRequest) (*ActivityResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ACTIVITY_SEGMENT, ACTIVITY_LATEST_ACLTIVITY_SEGMENT},
		RequestParams{
			"limit":           strconv.Itoa(params.Limit),
			"namespaces":      intArrToStr(params.Namespaces),
			"allowDuplicates": strconv.FormatBool(params.AllowDuplicates),
		})
	if err != nil {
		return nil, err
	}
	var result *ActivityResult
	return result, json.Unmarshal(jsonBlob, &result)
}

// Get recently changed articles
// http://muppet.wikia.com/api/v1#!/Activity/getRecentlyChangedArticles_get_1
func (wa *WikiaApi) ActivityRecentlyChangedArticles(params ActivityRequest) (*ActivityResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ACTIVITY_SEGMENT, ACTIVITY_RECENTLY_CHANGED_ARTICLES_SEGMENT},
		RequestParams{
			"limit":           strconv.Itoa(params.Limit),
			"namespaces":      intArrToStr(params.Namespaces),
			"allowDuplicates": strconv.FormatBool(params.AllowDuplicates),
		})
	if err != nil {
		return nil, err
	}
	var result *ActivityResult
	return result, json.Unmarshal(jsonBlob, &result)
}
