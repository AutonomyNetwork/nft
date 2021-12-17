package types

import (
	"github.com/google/uuid"
	"strings"
)

func GenUniqueID(prefix string) string {
	return prefix + strings.ReplaceAll(uuid.New().String(), "-", "")
}