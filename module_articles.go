package gowikia

import (
	"encoding/json"
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

type ArticleAsSimpleJsonResult struct {
	Sections AsSimpleJsonSections `json:"sections"`
}

// Get simplified article contents
// http://muppet.wikia.com/api/v1#!/Articles/getAsSimpleJson_get_0
func (wa *WikiaApi) ArticlesAsSimpleJson(id int) (ArticleAsSimpleJsonResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "AsSimpleJson"},
		RequestParams{
			"id": intToStr(id),
		})
	if err != nil {
		return ArticleAsSimpleJsonResult{}, err
	}
	var result ArticleAsSimpleJsonResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type DetailsResultItemRevision struct {
	Id        int    `json:"id"`
	User      string `json:"user"`
	UserId    int    `json:"user_id"`
	Timestamp int    `json:"timestamp"`
}

type DetailsResultItem struct {
	Id                 int                       `json:"id"`
	Title              string                    `json:"title"`
	Ns                 int                       `json:"ns"`
	Url                string                    `json:"url"`
	Revision           DetailsResultItemRevision `json:"revision"`
	Comments           int                       `json:"comments"`
	Type               string                    `json:"type"`
	Abstract           string                    `json:"abstract"`
	Thumbnail          string                    `json:"thumbnail"`
	OriginalDimensions OriginalDimensions        `json:"original_dimensions"`
}

type ArticleDetailsResult struct {
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

// Get details about one or more articles
// http://muppet.wikia.com/api/v1#!/Articles/getDetails_get_1
func (wa *WikiaApi) ArticlesDetails(params DetailsRequest) (ArticleDetailsResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Details"},
		RequestParams{
			"ids":      intArrToStr(params.ids),
			"titles":   strArrToStr(params.titles),
			"abstract": intToStr(params.abstract),
			"width":    intToStr(params.width),
			"height":   intToStr(params.height),
		})
	if err != nil {
		return ArticleDetailsResult{}, err
	}
	var result ArticleDetailsResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ListResultItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Ns    int    `json:"ns"`
}

type ListResult struct {
	Items    []ListResultItem `json:"items"`
	Offset   string           `json:"offset"`
	Basepath string           `json:"basepath"`
}

type ListRequest struct {
	category   string
	namespaces []int
	limit      int
	offset     string
}

// Get a list of pages on the current wiki
// http://muppet.wikia.com/api/v1#!/Articles/getList_get_3
func (wa *WikiaApi) ArticlesList(params ListRequest) (ListResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "List"},
		RequestParams{
			"category":   params.category,
			"namespaces": intArrToStr(params.namespaces),
			"limit":      intToStr(params.limit),
			"offset":     params.offset,
		})
	if err != nil {
		return ListResult{}, err
	}
	var result ListResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type MostLinkedResultItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Ns          int    `json:"ns"`
	BackLinkCnt int    `json:"backlink_cnt"`
}

type MostLinkedResult struct {
	Items    []MostLinkedResultItem `json:"items"`
	Basepath string                 `json:"basepath"`
}

// Get the most linked articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getTop_get_4
func (wa *WikiaApi) ArticlesMostLinked() (MostLinkedResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "MostLinked"},
		RequestParams{})
	if err != nil {
		return MostLinkedResult{}, err
	}
	var result MostLinkedResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type Creator struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type ArticlesNewResult struct {
	Id                 int                `json:"id"`
	Ns                 int                `json:"ns"`
	Title              string             `json:"title"`
	Abstract           string             `json:"abstract"`
	Quality            int                `json:"quality"`
	Url                string             `json:"url"`
	Creator            Creator            `json:"creator"`
	CreationDate       string             `json:"creation_date"` // WTF: String
	Thumbnail          string             `json:"thumbnail"`
	OriginalDimensions OriginalDimensions `json:"original_dimensions"`
}

type ArticlesNewRequest struct {
	Namespaces        []int
	Limit             int
	MinArticleQuality int
}

// Get list of new articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getNew_get_6
func (wa *WikiaApi) ArticlesNew(params ArticlesNewRequest) (ArticlesNewResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "New"},
		RequestParams{
			"namespaces":        intArrToStr(params.Namespaces),
			"limit":             intToStr(params.Limit),
			"minArticleQuality": intToStr(params.MinArticleQuality),
		})
	if err != nil {
		return ArticlesNewResult{}, err
	}
	var result ArticlesNewResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesPopularItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ArticlesPopularResult struct {
	Items    []ArticlesPopularItem `json:"items"`
	Basepath string                `json:"basepath"`
}

// Get popular articles for the current wiki (from the beginning of time)
// http://muppet.wikia.com/api/v1#!/Articles/getPopular_get_7
func (wa *WikiaApi) ArticlesPopular(limit int) (ArticlesPopularResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Popular"},
		RequestParams{
			"limit": intToStr(limit),
		})
	if err != nil {
		return ArticlesPopularResult{}, err
	}
	var result ArticlesPopularResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesTopItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Ns    int    `json:"ns"`
}

type ArticlesTopResult struct {
	Items    []ArticlesTopItem `json:"items"`
	Basepath string            `json:"basepath"`
}

type ArticlesTopRequest struct {
	Namespaces []int
	Category   string
	Limit      int // WTF: String in original doc
}

// Get the most viewed articles on this wiki
// http://muppet.wikia.com/api/v1#!/Articles/getTop_get_9
func (wa *WikiaApi) ArticlesTop(params ArticlesTopRequest) (ArticlesTopResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Top"},
		RequestParams{
			"namespaces": intArrToStr(params.Namespaces),
			"category":   params.Category,
			"limit":      intToStr(params.Limit),
		})
	if err != nil {
		return ArticlesTopResult{}, err
	}
	var result ArticlesTopResult
	return result, json.Unmarshal(jsonBlob, &result)
}

type ArticlesTopByHubWiki struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Domain   string `json:"domain"`
}

type ArticlesTopByHubArticle struct {
	Id int `json:"id"`
	Ns int `json:"ns"`
}

type ArticlesTopByHubItem struct {
	Wiki     ArticlesTopByHubWiki      `json:"wiki"`
	Articles []ArticlesTopByHubArticle `json:"articles"`
}

type ArticlesTopByHubResult struct {
	Items []ArticlesTopByHubItem `json:"items"`
}

type TopByHubRequest struct {
	Hub        string
	Lang       string
	Namespaces []int
}

// Get the top articles by pageviews for a hub
// http://muppet.wikia.com/api/v1#!/Articles/getTopByHub_get_11
func (wa *WikiaApi) ArticlesTopByHub(params TopByHubRequest) (ArticlesTopByHubResult, error) {
	jsonBlob, err := getJsonBlob(
		wa.url,
		[]string{ARTICLES_SEGMENT, "Top"},
		RequestParams{
			"hub":        params.Hub,
			"lang":       params.Lang,
			"namespaces": intArrToStr(params.Namespaces),
		})
	if err != nil {
		return ArticlesTopByHubResult{}, err
	}
	var result ArticlesTopByHubResult
	return result, json.Unmarshal(jsonBlob, &result)
}
