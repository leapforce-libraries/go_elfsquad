package elfsquad

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	Error struct {
		Code       string   `json:"code"`
		Message    string   `json:"message"`
		Details    []string `json:"details"`
		InnerError struct {
			Message    string `json:"message"`
			Type       string `json:"type"`
			Stacktrace string `json:"stacktrace"`
		}
	} `json:"error"`
}
