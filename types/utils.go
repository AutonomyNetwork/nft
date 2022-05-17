package types

import (
	"strings"
	
	"github.com/google/uuid"
)

func GenUniqueID(prefix string) string {
	return prefix + strings.ReplaceAll(uuid.New().String(), "-", "")
}
