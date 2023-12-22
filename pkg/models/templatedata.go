package models

// TemplateData is a type that represents the data to be passed to a template for rendering.
// It contains several maps for storing string, integer, and float values identified by string keys.
// The Data field is a map[string]any, allowing any type of data to be stored.
// The CSRFToken field holds the CSRF token for protecting against cross-site request forgery.
// The Flash field is used for displaying a flash message to the user.
// The Warning and Error fields are used to display warning and error messages respectively.
//
// Example usage:
//
//	data := TemplateData{
//	    StringMap: map[string]string{
//	        "key1": "value1",
//	    },
//	    IntMap: map[string]int{
//	        "key2": 42,
//	    },
//	    FloatMap: map[string]float32{
//	        "key3": 3.14,
//	    },
//	    Data: map[string]any{
//	        "key4": struct{ Name string }{"John"},
//	    },
//	    CSRFToken: "somerandomtoken",
//	    Flash:     "
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
