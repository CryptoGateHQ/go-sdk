package cryptogate

import "fmt"

// APIError is returned when the CryptoGate API responds with a non-2xx status.
type APIError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("cryptogate: %s (%s)", e.Message, e.Code)
	}
	return fmt.Sprintf("cryptogate: %s", e.Message)
}

// IsNotFound returns true if the error is a 404.
func IsNotFound(err error) bool {
	if e, ok := err.(*APIError); ok {
		return e.Status == 404
	}
	return false
}

// IsAuthError returns true if the error is a 401.
func IsAuthError(err error) bool {
	if e, ok := err.(*APIError); ok {
		return e.Status == 401
	}
	return false
}

// IsValidationError returns true if the error is a 400.
func IsValidationError(err error) bool {
	if e, ok := err.(*APIError); ok {
		return e.Status == 400
	}
	return false
}
