package main

import (
	"regexp"
)

func searchBL(bl []string, key string) bool {
	for _, x := range bl {
		if x == key {
			return true
		}
	}
	return false
}

func isSameDomain(siteUrl string, domain string) bool {
	reg := regexp.MustCompile(`https*://`)
	siteUrl = reg.ReplaceAllString(siteUrl, "${1}")

	ret, err := regexp.Match(`^(https*://)*[^/\:]*`+siteUrl, []byte(domain))
	if err != nil || !ret {
		return false
	}
	return true
}

func finalUrl(addr string, url string) (string, error) {
	ret, err := regexp.Match(`^[/(https*://)].+`, []byte(addr))
	if err != nil || !ret {
		return "", err
	}

	final := addr
	if addr[0] == '/' {
		final = urlPostfix(url, addr)
	}
	return final, nil
}

func regEx(pattern string, value string) []string {
	re := regexp.MustCompile(pattern)
	result := make([]string, 0)
	matches := re.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, match[1])
		} else {
			result = append(result, match[0])
		}
	}
	return result
}

func urlPostfix(url string, addr string) string {
	if url[len(url)-1] == '/' && addr[0] == '/' {
		url = url[:len(url)-1]
	}
	if url[len(url)-1] != '/' && addr[0] != '/' {
		url = url + "/"
	}
	return url + addr
}
