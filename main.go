package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

type imageCrawler struct {
	url     string
	ImgUrls map[string][]string
	imgBL   []string
	body    string
}

var wg sync.WaitGroup

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

func isSameDomain(siteUrl string, domain string) bool {
	ret, err := regexp.Match(`^(https*://)*[^/\:]*`+siteUrl, []byte(domain))
	if err != nil || !ret {
		return false
	}
	return true
}

func (img *imageCrawler) scanForImages(url string) error {
	pattern := `<img[^>]*src=["']([^>]*?)["'][^>]*>`
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	results := regEx(pattern, string(body[:]))
	imgs := make([]string, 0)
	for _, result := range results {
		final, err := finalUrl(result, img.url)
		if final == "" {
			final = result
		}
		if err != nil {
			return err
		}

		imgs = append(imgs, final)
	}

	img.ImgUrls[url] = imgs
	return nil
}

func (img *imageCrawler) crawlPack(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	paths := regEx(`<a[^>]*href=["']([^>]*?)["'][^>]*>`, string(body[:]))
	result := make([]string, 0)
	for _, path := range paths {
		final, _ := finalUrl(path, img.url)
		if final != "" {
			result = append(result, final)
		}
	}

	return result, nil
}

func searchBL(bl []string, key string) bool {
	for _, x := range bl {
		if x == key {
			return true
		}
	}
	return false
}

func (img *imageCrawler) crawl() error {
	queue := make([]string, 0)
	bl := make([]string, 0)
	queue = append(queue, img.url)

	for len(queue) > 0 {
		url := queue[0]
		queue = queue[1:]
		wg.Add(1)
		go func() {
			defer wg.Done()
			pack, _ := img.crawlPack(url)

			if !searchBL(bl, url) {
				img.scanForImages(url)
				bl = append(bl, url)
			}

			for _, x := range pack {
				if !searchBL(bl, x) {
					if isSameDomain(img.url, x) {
						queue = append(queue, x)
					}
				}
			}
		}()

		if len(queue) == 0 {
			wg.Wait()
		}
	}

	return nil
}

func initCrawler(url string) *imageCrawler {
	result := &imageCrawler{
		url:     url,
		ImgUrls: make(map[string][]string, 0),
	}
	return result
}

func main() {
	crawler := initCrawler("https://www.yartu.io/")
	err := crawler.crawl()
	if err != nil {
		panic(err)
	}

	wg.Wait()

	counter := 0
	for key, values := range crawler.ImgUrls {
		fmt.Printf("%s:\n", key)
		for i, value := range values {
			fmt.Printf("\t%d.) %s\n", i+1, value)
			counter++
		}
	}
	fmt.Printf("\nCOUNTER: %d\n", counter)
}
