package response

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Response is the standard JSON body for responses.
//
// Use Success is flag for success responses.
//
// Data is the ANY type payload for the response.
//
// Error is the ResponseError struct.
type Response struct {
	Success bool           `json:"success" example:"true"`
	Data    any            `json:"data,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
	// Can be added metadata to the response
}

type ResponseError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Validation failed"`
	Details any    `json:"details,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
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
