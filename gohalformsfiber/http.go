package gohalformsfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sazzer/gohalforms"
)

// Send sends a HAL (Hypertext Application Language) resource as a Fiber response to the client.
//
// Parameters:
//
//	c - The *fiber.Ctx instance representing the Fiber context to which the response will be sent.
//	resource - The gohalforms.Resource instance representing the HAL resource to be sent as a response.
//
// Returns:
//
//	An error if there was an issue sending the response; otherwise, it returns nil.
//
// Example:
//
//	// Create a new HAL resource.
//	halResource := gohalforms.New(map[string]any{
//	    "property1": "value1",
//	})
//
//	// Send the HAL resource as a Fiber response.
//	err := gohalformsfiber.Send(c, halResource)
//	if err != nil {
//	    // Handle the error, e.g., log it or send an alternative response.
//	}
func Send(c *fiber.Ctx, resource gohalforms.Resource) error {
	err := c.JSON(resource)
	if err != nil {
		return err
	}

	c.Response().Header.Set("Content-Type", resource.GetContentType())

	return nil
}
