package validation

// Request defines the input structure received by a http handler
type Request struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	URL   string `json:"url" validate:"url"`
}
