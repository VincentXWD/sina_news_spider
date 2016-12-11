package main

import (
	log "github.com/Sirupsen/logrus"
)

const (
	MAX_SPECIALCOVERAGE_SIZE = 6
	MAX_NEWS_SIZE = 1024
	MAX_SUBJECT_SIZE = 256
	SIGNAL = "_KIRAI_YOUSHOULDEXIT_KIRAI_"
	KIRAI_DEBUG = 2
	RESULT_PATH = "./result/"
		CATALOG_PATH = "./catalog/"
	SUBJECT_PATH_CAT_FILENAME = "subject.usns"
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

func Init() {
	CreateDir(RESULT_PATH)
	CreateDir(CATALOG_PATH)
	UpdateSubjectUrl()
}

func main() {
	log.Infoln("Start.")
	Init()
	urls := GetSubjectUrl()
	for i := 0; i < len(urls); i++ {
		for j := 0; j < len(urls[i]); j++ {
			GetNewsUrls(urls[i][j], specialCoverage[i], "./result/")
		}
	}
}