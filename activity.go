package gowikia

import (
	"encoding/json"
)

type LatestActivityResultItem struct {
	Article    int `json:"article"`
	User       int `json:"user"`
	RevisionId int `json:"revisionId"`
	Timestamp  int `json:"timestamp"`
}

type LatestActivityResult struct {
	Items    []LatestActivityResultItem `json:"items"`
	Basepath string                     `json:"basepath"`
}

func parseLatesActivity(jsonBlob []byte) (LatestActivityResult, error) {
	var result LatestActivityResult
	err := json.Unmarshal(jsonBlob, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (wa *WikiaApi) LatestActivity(limit int, namespaces []int, allowDuplicates bool) (LatestActivityResult, error) {
	path := generatePath([]string{"Activity", "getLatestActivity"})
	url := generateApiUrl(wa.url, path, "")
	jsonBlob, err := wa.fetchUrl(url)
	if err != nil {
		return LatestActivityResult{}, err
	}
	return parseLatesActivity(jsonBlob)
}
