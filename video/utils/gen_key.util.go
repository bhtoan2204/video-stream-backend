package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

// TODO: Enhance this function in future to generate a key for a file
func GenFileKey(fileName string, folder string) string {
	ext := filepath.Ext(fileName)
	baseName := fileName[:len(fileName)-len(ext)]
	timestamp := time.Now().Format("20060102-150405")
	if folder != "" {
		return fmt.Sprintf("%s/%s-%s%s", folder, baseName, timestamp, ext)
	}
	return fmt.Sprintf("%s-%s%s", baseName, timestamp, ext)
}
