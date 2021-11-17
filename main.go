package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func finalUrl(url string, addr string) (string, error) {
	ret, err := regexp.Match(`^/.*`, []byte(addr))
	if err != nil && ret {
		return "", err
	}
	if url[len(url)-1] == '/' && addr[0] == '/' {
		url = url[:len(url)-1]
	}
	if url[len(url)-1] != '/' && addr[0] != '/' {
		url = url + "/"
	}
	return url+addr, nil
}

func scanForImages(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`<img[^>]*src=["'](.*?)["'][^>]*>`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	result := make([]string, 0)
	for _, match := range matches {
		final, err := finalUrl(url, match[1])
		if err != nil {
			return nil, err
		}
		result = append(result, final)
	}
	return result, nil
}

func main(){
	url := "https://blog.logrocket.com/"
	a, err := scanForImages(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", a)
}
