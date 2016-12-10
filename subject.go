package main

import (
	"regexp"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
)

var subjectUrlRegExp = regexp.MustCompile(`http://[^\s]+/z/[^\s]+.shtml`)

// 获取专题的url，专题名字不在这里处理，但是直接提取<TITLE>没有问题。保存至specialCoverage下属，返回所有路径
func GetSubject(Url string, specialCoverage string) string {
	rawHtml := GetHtml(Url)
	ret := subjectUrlRegExp.FindAllString(rawHtml, -1)
	subjects := make([]string, MAX_SUBJECT_SIZE)
	var s_it int = 0
	for _, href := range ret {
		subjects[s_it] = href
		if KIRAI_DEBUG == 2 {
			log.Infoln(subjects[s_it])
		}
		s_it++
	}
	var savePath string = ""
	var str string = ""
	for it := 0; it < s_it; it++ {
		str += subjects[it]
		str += "\n"
	}
	path := "./result/"+specialCoverage+"/"
	SaveFile(path, "suburls.usns", str)
	savePath += path + "suburls.usns"
	savePath += "\n"
	return savePath
}

func UpdateSubjectUrl() {
	var subjectPath string = ""
	for it := 0; it < MAX_SPECIALCOVERAGE_SIZE; it++ {
		CreateDir(RESULT_PATH+specialCoverage[it])
		subjectPath += GetSubject(rootUrls[it], specialCoverage[it])
	}
	SaveFile(CATALOG_PATH, "subject.usns", subjectPath)
}

func GetSubjectUrl() [][]string {
	buf, err := ioutil.ReadFile(CATALOG_PATH+SUBJECT_PATH_CAT_FILENAME)
	if err != nil {
		log.Errorln(err.Error())
		return nil
	}
	paths := strings.Split(string(buf), "\n")
	var urls [][]string = make([][]string, MAX_SPECIALCOVERAGE_SIZE)
	for i := 0; i < len(paths); i++ {
		if len(paths[i]) == 0 {
			continue
		}
		buf, err = ioutil.ReadFile(paths[i])
		if err != nil {
			log.Errorln(err.Error())
			return nil
		}
		url := strings.Split(string(buf), "\n")
		urls[i] = url
	}
	return urls
}