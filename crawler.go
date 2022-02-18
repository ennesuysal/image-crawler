package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type imageCrawler struct {
	url     string
	ImgUrls map[string][]string
	body    string
}

var imgMutex = sync.RWMutex{}

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

	imgMutex.Lock()
	img.ImgUrls[url] = imgs
	imgMutex.Unlock()
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
