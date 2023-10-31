package gohalforms_test

import (
	"encoding/json"
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
