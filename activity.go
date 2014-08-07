package gowikia

import (
	"encoding/json"
	"strconv"
)

const (
	ACTIVITY_SEGMENT = "Activity"
)

type ActivityRequest struct {
	limit           int
	namespaces      []int
	allowDuplicates bool
}

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

// Provides default values for LatestActivity function
func ActivityDefaults() ActivityRequest {
	return ActivityRequest{10, []int{0}, true}
}

// Get latest activity information
// http://muppet.wikia.com/api/v1#!/Activity/getLatestActivity_get_0
func (wa *WikiaApi) LatestActivity(ar ActivityRequest) (ActivityResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ACTIVITY_SEGMENT, "getLatestActivity"},
		RequestParams{
			"limit":           strconv.Itoa(ar.limit),
			"namespaces":      intArrToStr(ar.namespaces),
			"allowDuplicates": strconv.FormatBool(ar.allowDuplicates),
		})
	if err != nil {
		return ActivityResult{}, err
	}
	var ares ActivityResult
	return ares, json.Unmarshal(jsonBlob, &ares)
}

// Get recently changed articles
// http://muppet.wikia.com/api/v1#!/Activity/getRecentlyChangedArticles_get_1
func (wa *WikiaApi) RecentlyChangedArticles(ar ActivityRequest) (ActivityResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ACTIVITY_SEGMENT, "RecentlyChangedArticles"},
		RequestParams{
			"limit":           strconv.Itoa(ar.limit),
			"namespaces":      intArrToStr(ar.namespaces),
			"allowDuplicates": strconv.FormatBool(ar.allowDuplicates),
		})
	if err != nil {
		return ActivityResult{}, err
	}
	var ares ActivityResult
	return ares, json.Unmarshal(jsonBlob, &ares)
}
