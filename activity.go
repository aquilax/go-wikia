package gowikia

import (
	"encoding/json"
)

type LatestActivityResultItem struct {
	Article    int
	User       int
	RevisionId int
	timestamp  int
}

type LatestActivityResult struct {
	Items    []LatestActivityResultItem
	Basepath string
}

func (wa *WikiaApi) LatestActivity(limit int, namespaces []int, allowDuplicates bool) (LatestActivityResult, error) {
	path := generatePath([]string{"Activity", "getLatestActivity"})
	url, err := generateApiUrl(wa.url, path, "")
	if err != nil {
		return LatestActivityResult{}, err
	}
	var jsonBlob []byte
	jsonBlob, err = wa.fetchUrl(url)
	if err != nil {
		return LatestActivityResult{}, err
	}
	var result LatestActivityResult
	err = json.Unmarshal(jsonBlob, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
