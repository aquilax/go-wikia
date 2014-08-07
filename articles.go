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

type DetailsResultItemRevision struct {
	Id        int    `json:"id"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	Timestamp int    `json:"timestamp"`
}

type DetailsResultItemOriginalDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type DetailsResultItem struct {
	Id                 int                                 `json:"id"`
	Title              string                              `json:"title"`
	Ns                 int                                 `json:"ns"`
	Url                string                              `json:"url"`
	Revision           DetailsResultItemRevision           `json:"revision"`
	Comments           int                                 `json:"comments"`
	Type               string                              `json:"type"`
	Abstract           string                              `json:"abstract"`
	Thumbnail          string                              `json:"thumbnail"`
	OriginalDimensions DetailsResultItemOriginalDimensions `json:"original_dimensions"`
}

type DetailsResult struct {
	Items    []DetailsResultItem `json:"items"`
	Basepath string              `json:"basepath"`
}

type DetailsRequest struct {
	ids      []int
	titles   []string
	abstract int
	width    int
	height   int
}

func DetailsDefaults() DetailsRequest {
	return DetailsRequest{[]int{50}, []string{}, 100, 200, 200}
}

func (wa *WikiaApi) Details(dr DetailsRequest) (DetailsResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Details"},
		RequestParams{
			"ids":      intArrToStr(dr.ids),
			"titles":   strArrToStr(dr.titles),
			"abstract": intToStr(dr.abstract),
			"width":    intToStr(dr.width),
			"height":   intToStr(dr.height),
		})
	if err != nil {
		return DetailsResult{}, err
	}
	var dres DetailsResult
	return dres, json.Unmarshal(jsonBlob, &dres)
}
