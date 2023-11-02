package gohalforms

import "encoding/json"

// Resource represents a generic representation of a HAL (Hypertext Application Language) resource.
type Resource struct {
	payload   any
	links     linkset
	embedded  resourceset
	templates map[string]Template
}

// New creates a new instance of the Resource type with the provided payload.
//
// Parameters:
//
//	payload - The payload associated with the HAL resource.
//
// Returns:
//
//	A Resource instance containing the specified payload.
//
// Example:
//
//	// Create a new HAL resource with a custom payload.
//	payload := map[string]any{
//	    "property1": "value1",
//	    "property2": "value2",
//	}
//	halResource := gohalforms.New(payload)
func NewResource(payload any) Resource {
	return Resource{
		payload:   payload,
		links:     linkset{},
		embedded:  resourceset{},
		templates: map[string]Template{},
	}
}

// AddLink adds a new hyperlink to the HAL (Hypertext Application Language) resource under the specified relation.
//
// Parameters:
//
//	resource - A pointer to the Resource instance to which the link should be added.
//	rel - The relation name under which the link will be stored.
//	value - The Link instance to be added.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Create a new link.
//	newLink := gohalforms.Link{
//	    Href: "https://example.com/resource",
//	    Title: "Example Resource",
//	}
//
//	// Add the link to the HAL resource under the "related" relation.
//	halResource.AddLink("related", newLink)
func (resource *Resource) AddLink(rel string, value Link) {
	resource.links[rel] = append(resource.links[rel], value)
}

// AddEmbedded adds a new embedded HAL resource to the HAL (Hypertext Application Language) resource under the specified relation.
//
// Parameters:
//
//	resource - A pointer to the Resource instance to which the embedded resource should be added.
//	rel - The relation name under which the embedded resource will be stored.
//	value - The Resource instance to be added as an embedded resource.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Create a new embedded resource.
//	embeddedResource := gohalforms.New(map[string]any{
//	    "embeddedProperty": "embeddedValue",
//	})
//
//	// Add the embedded resource to the HAL resource under the "items" relation.
//	halResource.AddEmbedded("items", embeddedResource)
func (resource *Resource) AddEmbedded(rel string, value Resource) {
	resource.embedded[rel] = append(resource.embedded[rel], value)
}

// AddTemplate adds a new template to the HAL (Hypertext Application Language) resource under the specified relation.
//
// Parameters:
//
//	resource - A pointer to the Resource instance to which the template should be added.
//	rel - The relation name under which the template will be stored.
//	value - The Template instance to be added as a template for creating or updating the resource.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Create a new template for creating or updating the resource.
//	newTemplate := gohalforms.Template{
//	    ContentType: "application/json",
//	    Method:      "POST",
//	    Target:      "https://example.com/resource",
//	    Title:       "Create Resource",
//	    Properties:  []Property{},
//	}
//
//	// Add the template to the HAL resource under the "create" relation.
//	halResource.AddTemplate("create", newTemplate)
func (resource *Resource) AddTemplate(rel string, value Template) {
	resource.templates[rel] = value
}

func (resource Resource) MarshalJSON() ([]byte, error) {
	intermediate := map[string]any{}

	if resource.payload != nil {
		// Marshal the payload portion to JSON.
		raw, err := json.Marshal(resource.payload)
		if err != nil {
			return nil, err
		}

		// Unmarshal this into a Map.
		if err = json.Unmarshal(raw, &intermediate); err != nil {
			return nil, err
		}
	}

	if len(resource.links) > 0 {
		intermediate["_links"] = resource.links
	}

	if len(resource.embedded) > 0 {
		intermediate["_embedded"] = resource.embedded
	}

	if len(resource.templates) > 0 {
		intermediate["_templates"] = resource.templates
	}

	// Re-marshal this to JSON.
	return json.Marshal(intermediate)
}
