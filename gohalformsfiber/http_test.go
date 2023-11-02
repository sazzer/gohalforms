package gohalformsfiber_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kinbiko/jsonassert"
	"github.com/sazzer/gohalforms"
	"github.com/sazzer/gohalforms/gohalformsfiber"
	"github.com/stretchr/testify/assert"
)

func TestSendEmptyResource(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		resource := gohalforms.NewResource(nil)

		return gohalformsfiber.Send(c, resource)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, []string{"application/json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{}`)
}

func TestSendPopulatedResource(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		resource := gohalforms.NewResource(map[string]any{
			"hello":  "World!",
			"answer": 42,
		})

		return gohalformsfiber.Send(c, resource)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, []string{"application/json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestSendResourceWithLinks(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		resource := gohalforms.NewResource(map[string]any{
			"hello":  "World!",
			"answer": 42,
		})
		resource.AddLink("self", gohalforms.Link{Href: "/testSelfLink"})

		return gohalformsfiber.Send(c, resource)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, []string{"application/hal+json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{
		"_links": {
			"self": {"href": "/testSelfLink"}
		},
		"hello":  "World!",
		"answer": 42
	}`)
}

func TestSendResourceWithEmbedded(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		resource := gohalforms.NewResource(map[string]any{
			"hello": "World!",
		})
		resource.AddEmbedded("other", gohalforms.NewResource(map[string]any{
			"age": 41,
		}))

		return gohalformsfiber.Send(c, resource)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, []string{"application/hal+json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{
		"_embedded": {
			"other": {
				"age":  41
			}
		},
		"hello":  "World!"
	}`)
}

func TestSendResourceWithTemplate(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		resource := gohalforms.NewResource(map[string]any{
			"hello": "World!",
		})
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

		return gohalformsfiber.Send(c, resource)
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	assert.NoError(t, err)

	defer response.Body.Close()

	assert.Equal(t, []string{"application/prs.hal-forms+json; charset=utf-8"}, response.Header.Values("content-type"))

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	ja := jsonassert.New(t)
	ja.Assertf(string(body), `{
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
		},
		"hello":  "World!"
	}`)
}
