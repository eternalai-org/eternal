package utils

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func MinifyHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		var str string
		str, err = item.Html()
		str = strings.TrimSpace(str)
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\t", "")
		if err == nil {
			item.SetHtml(str)
		}
	})

	htmlStr, err := doc.Html()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(htmlStr), nil
}

func ResolveURL(base, ref string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	refURL, err := url.Parse(ref)
	if err != nil {
		return "", err
	}
	return baseURL.ResolveReference(refURL).String(), nil
}

func IsBase64DataURL(urlStr string) bool {
	// Check if the URL starts with the data URL scheme
	if !strings.HasPrefix(urlStr, "data:") {
		return false
	}

	// Regular expression to match Base64 data URLs
	base64Pattern := `^data:[\w/]+;base64,[A-Za-z0-9+/=]+$`
	matched, err := regexp.MatchString(base64Pattern, urlStr)
	if err != nil {
		return false
	}

	return matched
}
