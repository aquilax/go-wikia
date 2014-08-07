package gowikia

import (
	"encoding/json"
)

const (
	ARTICLES_SEGMENT = "Articles"
)


type AsSimpleJsonElementsItem struct {
	Text     string                     `json:"text"`
	Elements []AsSimpleJsonElementsItem `json:"elements"`
}

type AsSimpleJsonContentItem struct {
	Type     string                     `json:"type"`
	Text     string                     `json:"text"`
	Elements []AsSimpleJsonElementsItem `json:"elements"`
}

type AsSimpleJsonImagesItem struct {
	Src     string `json:"src"`
	Caption string `json:"caption"`
}

type AsSimpleJsonSections struct {
	Title   string                    `json:"title"`
	Level   int                       `json:"level"`
	Content []AsSimpleJsonContentItem `json:"content"`
	Images  []AsSimpleJsonImagesItem  `json:"images"`
}

type AsSimpleJsonResult struct {
	Sections AsSimpleJsonSections `json:"sections"`
}

// Get simplified article contents
// http://muppet.wikia.com/api/v1#!/Articles/getAsSimpleJson_get_0
func (wa *WikiaApi) AsSimpleJson(id int) (AsSimpleJsonResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "AsSimpleJson"},
		RequestParams{
			"id": intToStr(id),
		})
	if err != nil {
		return AsSimpleJsonResult{}, err
	}
	var ares AsSimpleJsonResult
	return ares, json.Unmarshal(jsonBlob, &ares)
}
