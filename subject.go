package main

import (
	"regexp"
	log "github.com/Sirupsen/logrus"
)

var subjectUrlRegExp = regexp.MustCompile(`http://[^\s]+/z/[^\s]+.shtml`)
var subjectNameRegExp = regexp.MustCompile(`<title>.+</title>`)

// 获取专题的url，专题名字不在这里处理，但是直接提取<TITLE>没有问题。保存至specialCoverage下属
func GetSubject(Url string, specialCoverage string) {
	rawHtml := GetHtml(Url)
	ret := subjectUrlRegExp.FindAllString(rawHtml, -1)
	subjects := make([]Subject, MAX_SUBJECT_SIZE)
	var s_it int = 0
	for _, href := range ret {
		subjects[s_it].Url = href
		log.Infoln(subjects[s_it])
		s_it++
	}
	if KIRAI_DEBUG == 2 {
		var str string = ""
		for it := 0; it < s_it; it++ {
			str += subjects[it].Url
			str += "\n"
		}
		SaveFile("./result/"+specialCoverage+"/", "suburls.usns", str)
	}
}