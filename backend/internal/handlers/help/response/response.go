package response

import (
	"encoding/json"
	"net/http"
)

// Response is the standard JSON body for responses.
//
// Use Success is flag for success responses.
// Data is the ANY type payload for the response.
// Error is the ResponseError struct.
type Response struct {
	Success bool           `json:"success"`
	Data    any            `json:"data,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
	// Can be added metadata to the response
}

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// ErrorResponse is the standard JSON body for error responses.
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

// JSON writes v as JSON with the given code code.
func JSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// Success writes a success response with a data payload.
func Success(w http.ResponseWriter, code int, data any) {
	JSON(w, code, Response{Success: true, Data: data})
}

// Error writes a standardized error response.
func Error(w http.ResponseWriter, code int, message string, details any) {
	JSON(w, code, Response{Success: false, Error: &ResponseError{Code: code, Message: message, Details: details}})
}

// Dummy response for success
func Ok(w http.ResponseWriter) {
	JSON(w, http.StatusOK, Response{Success: true})
}
