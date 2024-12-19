package utils

import (
	"net/url"
	"strings"
)

func ExtractDomainFromUrl(link string) (string, error) {
	parser, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(parser.Hostname(), "www."), nil
}

// removeRootPath removes the specified root directory from the given file path.
func RemoveRootPath(path, root string) string {
	// Ensure the root is correctly formatted with a trailing slash
	root = strings.TrimRight(root, "/") + "/"

	// Remove the root directory from the path if it exists
	if strings.HasPrefix(path, root) {
		return strings.TrimPrefix(path, root)
	}
	return path
}

func ShouldDownloadFile(contentType string) bool {
	if contentType == "" ||
		strings.Contains(contentType, "text/html") ||
		strings.Contains(contentType, "video") {
		return false
	}
	return true
}
