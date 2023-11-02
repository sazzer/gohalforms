package gohalforms

import (
	"encoding/json"
	"net/http"
)

// GetContentType returns the appropriate content type for the HAL (Hypertext Application Language) resource based on its content.
//
// Returns:
//
//	A string representing the content type of the HAL resource. The content type is suitable for setting the "content-type" header in an HTTP response.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Get the content type for the HAL resource.
//	contentType := halResource.GetContentType()
func (resource Resource) GetContentType() string {
	if len(resource.templates) > 0 {
		return "application/prs.hal-forms+json; charset=utf-8"
	} else if len(resource.links) > 0 || len(resource.embedded) > 0 {
		return "application/hal+json; charset=utf-8"
	} else {
		return "application/json; charset=utf-8"
	}
}

// Send sends a HAL (Hypertext Application Language) resource as an HTTP response to the client.
//
// Parameters:
//
//	w - The http.ResponseWriter where the response will be written.
//	resource - The Resource instance representing the HAL resource to be sent as a response.
//
// Returns:
//
//	An error if there was an issue encoding and sending the response; otherwise, it returns nil.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Send the HAL resource as an HTTP response.
//	err := gohalforms.Send(w, halResource)
//	if err != nil {
//	    // Handle the error, e.g., log it or send an alternative response.
//	}
func Send(w http.ResponseWriter, resource Resource) error {
	w.Header().Add("content-type", resource.GetContentType())

	return json.NewEncoder(w).Encode(resource)
}
