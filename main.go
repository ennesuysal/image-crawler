package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type imageCrawler struct {
	url     string
	ImgUrls []string
	body    string
}

func regEx(pattern string, value string) []string {
	re := regexp.MustCompile(pattern)
	result := make([]string, 0)
	matches := re.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		result = append(result, match[1])
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

func (img *imageCrawler) scanForImages() error {
	pattern := `<img[^>]*src=["']([^>]*?)["'][^>]*>`
	results := regEx(pattern, img.body)
	for _, result := range results {
		final, err := finalUrl(result, img.url)
		if final == "" {
			final = result
		}
		if err != nil {
			return err
		}
		img.ImgUrls = append(img.ImgUrls, final)
	}
	return nil
}

func crawlPack(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	paths := regEx(`<a[^>]*href=["']([^>]*?)["'][^>]*>`, string(body[:]))
	result := make([]string, 0)
	for _, path := range paths {
		final, _ := finalUrl(path, url)
		if final != "" {
			result = append(result, final)
		}
	}

	return result, nil
}

func (img *imageCrawler) crawl() error {
	queue := make([]string, 0)
	queue = append(queue, img.url)

	pack, err := crawlPack(img.url)
	if err != nil {
		return err
	}

	// TODO RUN QUEUE
	fmt.Printf("%q", pack)

	return nil
}

func initCrawler(url string) *imageCrawler {
	result := &imageCrawler{
		url:     url,
		ImgUrls: make([]string, 0),
	}
	return result
}

func main() {
	crawler := initCrawler("https://www.yartu.io/")
	err := crawler.crawl()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", crawler.ImgUrls)
}
