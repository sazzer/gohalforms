package gohalforms_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/sazzer/gohalforms"
	"github.com/stretchr/testify/assert"
)

func TestMarshalEmpty(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(nil)

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{}`)
}

func TestMarshalMapBody(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello":  "World!",
		"answer": 42,
	})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestMarshalStructBody(t *testing.T) {
	t.Parallel()

	type body struct {
		Hello  string `json:"hello"`
		Answer int    `json:"answer"`
	}

	resource := gohalforms.NewResource(body{
		Hello:  "World!",
		Answer: 42,
	})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestMarshalIntegerBody(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(1)

	_, err := json.Marshal(resource)
	assert.Error(t, err)
}

func TestMarshalSelfLink(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello":  "World!",
		"answer": 42,
	})
	resource.AddLink("self", gohalforms.Link{Href: "/testSelfLink"})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_links": {
			"self": {"href": "/testSelfLink"}
		},
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestMarshalTwoLinksDifferentRels(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello":  "World!",
		"answer": 42,
	})
	resource.AddLink("a", gohalforms.Link{Href: "/a/b"})
	resource.AddLink("b", gohalforms.Link{Href: "/c/d"})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_links": {
			"a": {"href": "/a/b"},
			"b": {"href": "/c/d"}
		},
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestMarshalTwoLinksSameRel(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello":  "World!",
		"answer": 42,
	})
	resource.AddLink("a", gohalforms.Link{Href: "/a/b", Title: "First"})
	resource.AddLink("a", gohalforms.Link{Href: "/c/d", Title: "Second"})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_links": {
			"a": [
				{"href": "/a/b", "title": "First"},
				{"href": "/c/d", "title": "Second"}
			]
		},
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestMarshalSingleEmbedded(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello": "World!",
	})
	resource.AddEmbedded("other", gohalforms.NewResource(map[string]any{
		"age": 41,
	}))

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_embedded": {
			"other": {
				"age":  41
			}
		},
		"hello":  "World!"
	}`)
}

func TestMarshalTwoEmbeddedDifferentRels(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello": "World!",
	})
	resource.AddEmbedded("other1", gohalforms.NewResource(map[string]any{
		"age": 41,
	}))
	resource.AddEmbedded("other2", gohalforms.NewResource(map[string]any{
		"answer": 42,
	}))

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_embedded": {
			"other1": {
				"age":  41
			},
			"other2": {
				"answer":  42
			}
		},
		"hello":  "World!"
	}`)
}

func TestMarshalTwoEmbeddedSameRel(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(map[string]any{
		"hello": "World!",
	})
	resource.AddEmbedded("other", gohalforms.NewResource(map[string]any{
		"age": 41,
	}))
	resource.AddEmbedded("other", gohalforms.NewResource(map[string]any{
		"answer": 42,
	}))

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_embedded": {
			"other": [{
				"age":  41
			}, {
				"answer":  42
			}]
		},
		"hello":  "World!"
	}`)
}

func TestMarshalSimpleTemplate(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(nil)
	resource.AddTemplate("default", gohalforms.Template{
		Title:       "Create",
		Method:      http.MethodPost,
		ContentType: "application/json",
		Properties: []gohalforms.Property{
			{
				Name:     "title",
				Required: true,
				Prompt:   "Title",
			},
			{
				Name:   "completed",
				Value:  "false",
				Prompt: "Completed",
			},
		},
	})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_templates" : {
			"default" : {
				"title" : "Create",
				"method" : "POST",
				"contentType" : "application/json",
				"properties" : [
					{"name" : "title", "required" : true, "prompt" : "Title"},
					{"name" : "completed", "value" : "false", "prompt" : "Completed"}
				]
			}
		}
	}`)
}

func TestMarshalLinkOptions(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(nil)
	resource.AddTemplate("default", gohalforms.Template{
		Properties: []gohalforms.Property{
			{
				Name: "options",
				Options: gohalforms.LinkOption{
					Link:           gohalforms.Link{Href: "/options"},
					MaxItems:       3,
					SelectedValues: []string{"a", "b", "c"},
				},
			},
		},
	})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_templates" : {
			"default" : {
				"properties" : [
					{
						"name" : "options",
						"options": {
							"link": {"href": "/options"},
							"maxItems": 3,
							"selectedValues": ["a", "b", "c"]
						}
					}
				]
			}
		}
	}`)
}

func TestMarshalInlineOptions(t *testing.T) {
	t.Parallel()

	resource := gohalforms.NewResource(nil)
	resource.AddTemplate("default", gohalforms.Template{
		Properties: []gohalforms.Property{
			{
				Name: "options",
				Options: gohalforms.InlineOption{
					Inline: []gohalforms.InlineOptionValue{
						{Prompt: "First", Value: "1"},
						{Prompt: "Second", Value: "2"},
						{Prompt: "Third", Value: "3"},
					},
					MaxItems:       3,
					SelectedValues: []string{"1", "2", "3"},
				},
			},
		},
	})

	encoded, err := json.Marshal(resource)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(encoded), `{
		"_templates" : {
			"default" : {
				"properties" : [
					{
						"name" : "options",
						"options": {
							"inline": [
								{"prompt": "First", "value": "1"},
								{"prompt": "Second", "value": "2"},
								{"prompt": "Third", "value": "3"}
							],
							"maxItems": 3,
							"selectedValues": ["1", "2", "3"]
						}
					}
				]
			}
		}
	}`)
}
