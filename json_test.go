package gowikia

import (
	"encoding/json"
	"testing"
)

type jsonTestCase struct {
	blob         []byte
	getStructure func() interface{}
}

func TestJsonStructures(t *testing.T) {
	testCases := []jsonTestCase{
		{
			[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4}],"basepath":"http://example.com"}`),
			func() interface{} { return ActivityResult{} },
		},
		{
			[]byte(`{"items":[{"article":1,"user":2,"revisionId":3,"timestamp":4},{"article":5,"user":6,"revisionId":7,"timestamp":8}],"basepath":"http://example.com"}`),
			func() interface{} { return ActivityResult{} },
		},
		{
			[]byte(`{"sections":{"title":"string","level":1,"content":[{"type":"string","text":"string","elements":[{"text":"string","elements":[]}]}],"images":[{"src":"string","caption":"string"}]}}`),
			func() interface{} { return ArticlesAsSimpleJsonResult{} },
		},

		{
			[]byte(`{"items":[{"id":1,"title":"string","ns":2,"url":"string","revision":{"id":3,"user":"string","user_id":4,"timestamp":5},"comments":6,"type":"string","abstract":"string","thumbnail":"string","original_dimensions":{"width":7,"height":8}}],"basepath":"string"}`),
			func() interface{} { return ArticlesDetailsResult{} },
		},
		{
			[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":1}],"offset":"string","basepath":"string"}`),
			func() interface{} { return ArticlesListResult{} },
		},
		{
			[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":2,"backlink_cnt":3}],"basepath":"string"}`),
			func() interface{} { return ArticlesMostLinkedResult{} },
		},
		{
			[]byte(`{"id":1,"ns":2,"title":"string","abstract":"string","quality":3,"url":"string","creator":{"avatar":"string","name":"string"},"creation_date":"string","thumbnail":"string","original_dimensions":{"width":5,"height":6}}`),
			func() interface{} { return ArticlesNewResult{} },
		},
		{
			[]byte(`{"items":[{"id":1,"title":"string","url":"string"}],"basepath":"string"}`),
			func() interface{} { return ArticlesPopularResult{} },
		},
		{
			[]byte(`{"items":[{"id":1,"title":"string","url":"string","ns":2}],"basepath":"string"}`),
			func() interface{} { return ArticlesTopResult{} },
		},
		{
			[]byte(`{"items":[{"wiki":{"id":1,"name":"string","language":"string","domain":"string"},"articles":[{"id":2,"ns":3}]}]}`),
			func() interface{} { return ArticlesTopByHubResult{} },
		},
		{
			[]byte(`{"navigation":{"wikia":[{"text":"string","href":"string","children":[{"text":"string","href":"string"}]}],"wiki":[]}}`),
			func() interface{} { return NavigationDataResult{} },
		},
		{
			[]byte(`{"items":[{"url":"string","title":"string","id":1,"imgUrl":"string","text":"string"}],"basepath":"string"}`),
			func() interface{} { return RelatedPagesListResult{} },
		},
		{
			[]byte(`{"total":"int","batches":"int","currentBatch":"int","next":"int","items":[{"id":"int","title":"string","url":"string","ns":"int","quality":"int"}]}`),
			func() interface{} { return SearchListResult{} },
		},
		{
			[]byte(`{"items":[{"title":"string"}]}`),
			func() interface{} { return SearchSuggestionsListResult{} },
		},
		{
			[]byte(`{"items":[{"user_id":"int","title":"string","name":"string","url":"string","numberofedits":"int","avatar":"string"}],"basepath":"string"}`),
			func() interface{} { return UserDetailsResult{} },
		},
	}
	for i, testCase := range testCases {
		structure := testCase.getStructure()
		_ = json.Unmarshal(testCase.blob, &structure)
		result, _ := json.Marshal(structure)
		if len(testCase.blob) != len(result) {
			t.Errorf("For testCase %d expected %s but got %s",
				i, testCase.blob, result)
		}
	}
}
