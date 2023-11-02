package gohalforms

// Template represents a template for creating or updating a HAL (Hypertext Application Language) resource.
type Template struct {
	ContentType string     `json:"contentType,omitempty"`
	Method      string     `json:"method,omitempty"`
	Target      string     `json:"target,omitempty"`
	Title       string     `json:"title,omitempty"`
	Properties  []Property `json:"properties"`
}

// Property represents a property definition within a HAL resource template.
type Property struct {
	Name        string         `json:"name"`
	Prompt      string         `json:"prompt,omitempty"`
	Readonly    bool           `json:"readOnly,omitempty"`
	Regex       string         `json:"regex,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Templated   bool           `json:"templated,omitempty"`
	Value       string         `json:"value,omitempty"`
	Cols        uint32         `json:"cols,omitempty"`
	Max         uint32         `json:"max,omitempty"`
	MaxLength   uint32         `json:"maxLength,omitempty"`
	Min         uint32         `json:"min,omitempty"`
	MinLength   uint32         `json:"minLength,omitempty"`
	Options     PropertyOption `json:"options,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
	Rows        uint32         `json:"rows,omitempty"`
	Step        uint32         `json:"step,omitempty"`
	Type        string         `json:"type,omitempty"`
}

// PropertyOption is an interface for representing various property options in a HAL resource template.
type PropertyOption interface {
	isAnOption()
}

// InlineOptionValue represents a single value within a set of inline options.
type InlineOptionValue struct {
	Prompt string `json:"prompt"`
	Value  string `json:"value"`
}

// InlineOption represents a property option for inline values within a HAL resource template.
type InlineOption struct {
	Inline         []InlineOptionValue `json:"inline"`
	MaxItems       uint32              `json:"maxItems,omitempty"`
	MinItems       uint32              `json:"minItems,omitempty"`
	SelectedValues []string            `json:"selectedValues,omitempty"`
}

// isAnOption is a method to indicate that InlineOption implements the PropertyOption interface.
func (InlineOption) isAnOption() {}

// LinkOption represents a property option for link values within a HAL resource template.
type LinkOption struct {
	Link           Link     `json:"link"`
	MaxItems       uint32   `json:"maxItems,omitempty"`
	MinItems       uint32   `json:"minItems,omitempty"`
	SelectedValues []string `json:"selectedValues,omitempty"`
}

// isAnOption is a method to indicate that LinkOption implements the PropertyOption interface.
func (LinkOption) isAnOption() {}
