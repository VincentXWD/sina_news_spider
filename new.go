package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"html"
	"strings"
)

var newsRegExp = regexp.MustCompile(`http://english.sina.com/[^\s]+/[\w]*/*[\d]*/[\d]*/[\d]*.html`)
var newidRegExp = regexp.MustCompile(`[\d]+.html`)
var snewidRegExp = regexp.MustCompile(`\d+`)
var subjectNameRegExp = regexp.MustCompile(`<title>.+</title>`)

// 获取newid，直接从url中截取
func GetNewId(Url string) string {
	ret := newidRegExp.FindAllString(Url, -1)
	if len(ret) == 1 {
		sret := snewidRegExp.FindAllString(ret[0], -1)
		if len(sret) == 1 {
			return sret[0]
		}
		return ""
	}
	return ""
}

//提取title中包含的专题名
func GetSubjectName(rawHtml string) string {
	tmp := subjectNameRegExp.FindAllString(rawHtml, -1)
	if len(tmp) != 1 {
		return ""
	}
	tmp[0] = html.UnescapeString(tmp[0][7:len(tmp[0])-8])
	var it int = 0
	for it = 0; it < len(tmp[0]); it++ {
		if tmp[0][it] == '_' {
			break
		}
	}
	return tmp[0][:it]
}

//TODO: 获取新闻的url，直接保存到对应专题的path
func GetNewsUrls(Url string, specialCoverage string, savePath string) {
	var DEBUG_ALL string = ""
	rawHtml := GetHtml(Url)
	subjectName := GetSubjectName(rawHtml)
	if subjectName == "" {
		log.Errorln("No such subject. Ignore...")
		return
	}
	log.Infoln(subjectName)

	ret := newsRegExp.FindAllString(rawHtml, -1)
	news := make([]New, MAX_NEWS_SIZE)
	var news_it int = 0
	subjectName = strings.Replace(subjectName, " ", "_", -1)

	for _, href := range ret {
		log.Println("NEW:",href)
		if KIRAI_DEBUG == 1 {
			DEBUG_ALL += href
			DEBUG_ALL += "\n"
		}
		new := GetNews(href, subjectName, specialCoverage)
		if new.Subject == SIGNAL {
			continue
		}
		news[news_it] = new
		news_it++
	}
	if KIRAI_DEBUG == 1 {
		SaveFile(savePath, "urls", DEBUG_ALL)
	}
	log.Println("Special Coverage : ", specialCoverage)
	log.Println("Total : ", len(ret))
	log.Println("Saving news. Path : ", savePath)

	for n_it := 0; n_it < news_it; n_it++ {
		CreateDir(savePath+news[n_it].Prefix+"/"+news[n_it].Subject+"/")
		SaveNews(news[n_it], savePath+news[n_it].Prefix+"/"+news[n_it].Subject)
	}
}

// 保存新闻，目录为对应专题，保存newid.sns
func SaveNews(new New, savePath string) {
	var fileName string = new.NewId + ".sns"
	var buf string = new.Subject + "\n" + new.Title + "\n" + new.Time + "\n" + new.Content + "\n"
	SaveFile(savePath+"/", fileName, buf)
}

// TODO: query获取新闻时间、标题、内容
func GetNews(Url string, Subject string, specialCoverage string) New {
	doc, err := goquery.NewDocument(Url)
	var new = New{"","","","","",""}
	if err != nil {
		log.Errorln(err.Error())
		return New{SIGNAL,"","","","",""}
	}
	new.NewId = GetNewId(Url)
	if new.NewId == "" {
		return New{SIGNAL,"","","","",""}
	}
	new.Prefix = specialCoverage
	new.Subject = Subject
	new.Title = doc.Find("#Esinawrap .Main #Article .Title h1").First().Text()
	new.Time = doc.Find("#Esinawrap .Main #Article .Title .attribute span").First().Text()
	contentSelection := doc.Find("#Esinawrap .Main #Article .Content")
	for s_it := 0; s_it < contentSelection.Size(); s_it++ {
		content := contentSelection.Get(s_it)
		new.Content += goquery.NewDocumentFromNode(content).Find("p").Text()
	}
	contentSelection = doc.Find("#Esinawrap .Main #Article .b_cont")
	for s_it := 0; s_it < contentSelection.Size(); s_it++ {
		content := contentSelection.Get(s_it)
		new.Content += goquery.NewDocumentFromNode(content).Find("p").Text()
	}

	return new
}