package gohalforms

import "encoding/json"

// Link represents a hyperlink within a HAL (Hypertext Application Language) resource.
type Link struct {
	Href        string `json:"href"`
	Templated   bool   `json:"templated,omitempty"`
	Type        string `json:"type,omitempty"`
	Deprecation string `json:"deprecation,omitempty"`
	Name        string `json:"name,omitempty"`
	Profile     string `json:"profile,omitempty"`
	Title       string `json:"title,omitempty"`
	HrefLang    string `json:"hreflang,omitempty"`
}

// links is a slice of Link instances used to represent multiple links within a HAL resource.
type links []Link

// linkset is a map where the keys represent relation names, and the values are slices of Link instances.
type linkset map[string]links

// MarshalJSON serializes a links slice to JSON. If there is only one Link in the slice, it is serialized individually.
//
// Returns:
//
//	A JSON representation of the links slice or a single Link instance.
func (links links) MarshalJSON() ([]byte, error) {
	if len(links) == 1 {
		return json.Marshal(links[0])
	}

	return json.Marshal([]Link(links))
}
