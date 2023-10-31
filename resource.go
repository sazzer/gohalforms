package gohalforms

import "encoding/json"

// Resource represents a generic representation of a HAL (Hypertext Application Language) resource.
type Resource struct {
	payload any
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
		payload: payload,
	}
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

	// Re-marshal this to JSON.
	return json.Marshal(intermediate)
}
