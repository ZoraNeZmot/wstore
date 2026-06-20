package handlershelp

import (
	"encoding/json"
	"net/http"
	"path"
)

type JsonResponseToken struct {
	Token string `json:"token"`
}

type JsonResponseError struct {
	Error string `json:"error"`
}
type JsonResponseMessage struct {
	Message string `json:"message"`
}

const (
	JsonOk = `{\"message\": \"ok\"}`
)

func JSONError(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JsonResponseError{
		Error: err,
	})
}

const (
	BasePathAPI = "/api"
)

func GetFullPathAPI(relativePath string) string {
	return path.Join(BasePathAPI, relativePath)
}
