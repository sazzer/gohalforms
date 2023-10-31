package gohalforms

import "encoding/json"

// Resource represents a generic representation of a HAL (Hypertext Application Language) resource.
type Resource struct {
	payload  any
	links    linkset
	embedded resourceset
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
		payload:  payload,
		links:    linkset{},
		embedded: resourceset{},
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

func (resource *Resource) AddEmbedded(rel string, value Resource) {
	resource.embedded[rel] = append(resource.embedded[rel], value)
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

	// Re-marshal this to JSON.
	return json.Marshal(intermediate)
}
