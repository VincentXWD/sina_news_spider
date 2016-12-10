package main

import (
	log "github.com/Sirupsen/logrus"
	"regexp"
)

const (
	MAX_NEWS_SIZE = 1024
	MAX_SUBJECT_SIZE = 256
	SIGNAL = "_KIRAI_YOUSHOULDEXIT_KIRAI_"
	KIRAI_DEBUG = 2
)

var rootUrls = [...]string {
	"http://english.sina.com/specialcoverage/china.html",
	"http://english.sina.com/specialcoverage/world.html",
	"http://english.sina.com/specialcoverage/ent.html",
	"http://english.sina.com/specialcoverage/sports.html",
	"http://english.sina.com/specialcoverage/biz-tech.html",
	"http://english.sina.com/specialcoverage/culture.html",
}

var specialCoverage = [...]string {
	"china",
	"world",
	"ent",
	"sports",
	"biz_tech",
	"culture",
}

//var urlChannel = make(chan string, 200)

func main() {
	CreateDir("./result/")
	for i := range rootUrls {
		CreateDir("./result/"+specialCoverage[i])
		//GetNewsUrls(rootUrls[i], specialCoverage[i], "./result/"+specialCoverage[i]+"/")
		GetSubject(rootUrls[i], specialCoverage[i])
		log.Infoln("\n\n\n")
	}
	//SaveFile("./","test", GetHtml("http://english.sina.com/specialcoverage/ent.html"))
}