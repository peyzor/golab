package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var cnt = SafeCounter{v: make(map[string]bool)}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, r chan string) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	cnt.mux.Lock()
	_, ok := cnt.v[url]
	if ok {
		cnt.mux.Unlock()
		return
	}

	cnt.v[url] = true
	cnt.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	var wg_ sync.WaitGroup
	for _, u := range urls {
		wg_.Add(1)
		go Crawl(u, depth-1, fetcher, &wg_, r)
	}
	wg_.Wait()
	return
}

func runCrawler() {
	result := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg, result)
	wg.Wait()
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/xd/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
