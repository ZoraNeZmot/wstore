package handlershelp

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"
	tokenpkg "wstore/pkg/token"

	"github.com/google/uuid"
)

func IfEmptyResult(err error) bool {
	if err == nil {
		return true
	}
	if strings.HasPrefix(err.Error(), "empty query value:") || strings.HasPrefix(err.Error(), "empty path value:") || strings.HasPrefix(err.Error(), "empty value:") {
		return true
	}
	return false
}

// formatPersonName returns name with the first letter uppercased and the rest lowercased.
func FormatPersonName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return name
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func GetIntFromQuery(r *http.Request, key string, defaultvalue ...int) (int, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		if len(defaultvalue) > 0 {
			return defaultvalue[0], nil
		}
		return 0, fmt.Errorf("empty query value: %s", key)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		if len(defaultvalue) > 0 {
			return defaultvalue[0], nil
		}
		return 0, fmt.Errorf("invalid query value: %s", value)
	}
	return intValue, nil
}

func GetStringFromPath(r *http.Request, key string) (string, error) {
	value := r.PathValue(key)
	if value == "" {
		return "", fmt.Errorf("empty path value: %s", key)
	}
	return value, nil
}

func GetStringFromQuery(r *http.Request, key string, defaultvalue ...string) (string, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		if len(defaultvalue) > 0 {
			return defaultvalue[0], nil
		}
		return "", fmt.Errorf("empty query value: %s", key)
	}
	return value, nil
}

func GetUUIDFromPath(r *http.Request, key string) (uuid.UUID, error) {
	value := r.PathValue(key)
	if value == "" {
		return uuid.Nil, fmt.Errorf("empty path value: %s", key)
	}
	return uuid.Parse(value)
}

func GetUUIDFromQuery(r *http.Request, key string) (uuid.UUID, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return uuid.Nil, fmt.Errorf("empty query value: %s", key)
	}
	return uuid.Parse(value)
}

func GetClaimsFromContext(r *http.Request) (*tokenpkg.TokenClaims, error) {
	claims, ok := r.Context().Value(tokenpkg.ClaimsKey).(*tokenpkg.TokenClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get claims from context")
	}
	return claims, nil
}
