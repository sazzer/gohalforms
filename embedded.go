package gohalforms

import "encoding/json"

// resources is a slice of Resource instances used to represent multiple embedded HAL (Hypertext Application Language) resources.
type resources []Resource

// resourceset is a map where the keys represent relation names, and the values are slices of links that point to HAL resources.
type resourceset map[string]resources

// MarshalJSON serializes a resources slice to JSON. If there is only one Resource in the slice, it is serialized individually.
//
// Returns:
//
//	A JSON representation of the resources slice or a single Resource instance.
func (resources resources) MarshalJSON() ([]byte, error) {
	if len(resources) == 1 {
		return json.Marshal(resources[0])
	}

	return json.Marshal([]Resource(resources))
}
