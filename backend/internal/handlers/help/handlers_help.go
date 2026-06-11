package handlershelp

import (
	"path"
)

const (
	BasePathAPI = "/api"
)

func GetFullPathAPI(relativePath string) string {
	return path.Join(BasePathAPI, relativePath)
}
