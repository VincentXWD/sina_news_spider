package main

import (
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var userAgent = [...]string {
	"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
}

func GetRandomUserAgent() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgent[r.Intn(len(userAgent))]
}

func GetFakeHeader(request *http.Request) {
	request.Header.Set("User-Agent", GetRandomUserAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	//request.Header.Set("Accept-Encoding", "gzip, deflate")
}

func GetByProxy(proxyAddr string, Url string) (*http.Response, error) {
	timeout := time.Duration(10 * time.Second)
	request, err := http.NewRequest(http.MethodGet, Url, nil)
	GetFakeHeader(request)
	if err != nil {
		return nil, err
	}
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Errorln(err.Error())
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}
	return client.Do(request)
}

func GetByDirectory(Url string) (*http.Response, error) {
	timeout := time.Duration(10 * time.Second)
	request, err := http.NewRequest(http.MethodGet, Url, nil)
	GetFakeHeader(request)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: timeout,
	}
	return client.Do(request)
}
