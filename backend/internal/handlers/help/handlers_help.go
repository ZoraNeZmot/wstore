package handlershelp

import (
	"path"
)

const (
	BasePathAPI = "/api/v1"
	BasePathWeb = "/"
)

func GetFullPathAPI(relativePath string) string {
	return path.Join(BasePathAPI, relativePath)
}

func GetFullPathWeb(relativePath string) string {
	return path.Join(BasePathWeb, relativePath)
}
