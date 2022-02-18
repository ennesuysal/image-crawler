package main

import (
	"testing"
)

func TestSearchBL(t *testing.T) {
	bl := []string{"val1", "val2", "val3"}

	returnVal := searchBL(bl, "val1")
	expectedVal := true
	if returnVal != expectedVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	returnVal = searchBL(bl, "notfound")
	expectedVal = false
	if returnVal != expectedVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}
}

func TestIsSameDomain(t *testing.T) {
	siteUrl := "https://example.com"
	domain := "example.com"
	expectedVal := true
	returnVal := isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "example.com"
	domain = "https://example.com"
	expectedVal = true
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "example.com"
	domain = "example.com"
	expectedVal = true
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "https://example.com"
	domain = "https://example.com"
	expectedVal = true
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "https://differentsite.com"
	domain = "example.com"
	expectedVal = false
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "differentsite.com"
	domain = "https://example.com"
	expectedVal = false
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "differentsite.com"
	domain = "example.com"
	expectedVal = false
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}

	siteUrl = "https://example.com"
	domain = "https://differentsite.com"
	expectedVal = false
	returnVal = isSameDomain(siteUrl, domain)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%t]\tExpected Value [%t]", returnVal, expectedVal)
	}
}
func TestFinalUrl(t *testing.T) {
	addr := "/article/12"
	url := "https://examplesite.org"
	expectedVal := "https://examplesite.org/article/12"
	returnVal, err := finalUrl(addr, url)

	if expectedVal != returnVal || err != nil {
		t.Errorf("Test Failed: Return Value [%s]\tExpected Value [%s]", returnVal, expectedVal)
	}

	addr = "https://differentsite.com/article/12"
	url = "https://examplesite.org"
	expectedVal = addr
	returnVal, err = finalUrl(addr, url)

	if expectedVal != returnVal || err != nil {
		t.Errorf("Test Failed: Return Value [%s]\tExpected Value [%s]", returnVal, expectedVal)
	}
}

func TestRegEx(t *testing.T) {
	pattern := `[\^ ](a[^ ]*)`
	value := "banana apricot strawberry apple"
	expectedVal := []string{"apricot", "apple"}
	returnVal := regEx(pattern, value)

	if len(expectedVal) != len(returnVal) {
		t.Errorf("Test Failed: Return Value [%v]\tExpected Value [%v]", returnVal, expectedVal)
		return
	}

	for i, x := range returnVal {
		if x != expectedVal[i] {
			t.Errorf("Test Failed: Return Value [%v]\tExpected Value [%v]", returnVal, expectedVal)
			return
		}
	}
}

func TestUrlPostfix(t *testing.T) {
	url := "https://blabla.site"
	path := "/article/17"
	expectedVal := "https://blabla.site/article/17"
	returnVal := urlPostfix(url, path)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%s]\tExpected Value [%s]", returnVal, expectedVal)
	}

	url = "https://blabla.site/"
	path = "/article/17"
	expectedVal = "https://blabla.site/article/17"
	returnVal = urlPostfix(url, path)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%s]\tExpected Value [%s]", returnVal, expectedVal)
	}

	url = "https://blabla.site"
	path = "article/17"
	expectedVal = "https://blabla.site/article/17"
	returnVal = urlPostfix(url, path)
	if expectedVal != returnVal {
		t.Errorf("Test Failed: Return Value [%s]\tExpected Value [%s]", returnVal, expectedVal)
	}
}
