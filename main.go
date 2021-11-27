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
}

func (img *imageCrawler) finalUrl(addr string) (string, error) {
	ret, err := regexp.Match(`^/.*`, []byte(addr))
	if err != nil || !ret {
		return "", err
	}
	url := img.url
	if url[len(url)-1] == '/' && addr[0] == '/' {
		url = url[:len(url)-1]
	}
	if url[len(url)-1] != '/' && addr[0] != '/' {
		url = url + "/"
	}
	return url + addr, nil
}

func (img *imageCrawler) scanForImages() error {
	url := img.url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`<img[^>]*src=["']([^>]*?)["'][^>]*>`)
	matches := re.FindAllStringSubmatch(string(body), -1)
	for _, match := range matches {
		final, err := img.finalUrl(match[1])
		if final == "" {
			final = match[1]
		}
		if err != nil {
			return err
		}
		img.ImgUrls = append(img.ImgUrls, final)
	}
	return nil
}

func initCrawler(url string) *imageCrawler {
	return &imageCrawler{
		url:     url,
		ImgUrls: make([]string, 0),
	}
}

func main() {
	crawler := initCrawler("https://www.yartu.io/")
	err := crawler.scanForImages()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", crawler.ImgUrls)
}
