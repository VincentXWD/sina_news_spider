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

//按照对应template爬取新闻
func GewNews(urls []string, subjectName string, specialCoverage string) ([]New, int){
	news := make([]New, MAX_NEWS_SIZE)
	var news_it int = 0
	for _, href := range urls {
		log.Infoln("NEW:",href)
		new := getNews(href, subjectName, specialCoverage)
		if new.Prefix == SIGNAL {
			continue
		}
		news[news_it] = new
		news_it++
	}
	return news, news_it
}
//保存新闻，同时保存对应目录
func SaveNews(news []New, news_it int, savePath string, specialCoverage string) {
	CreateDir(CATALOG_PATH+specialCoverage+"/")
	if PathExists(CATALOG_PATH+specialCoverage+"/"+"path.snp") == false {
		SaveFile(CATALOG_PATH+specialCoverage+"/", "path.snp", "")
	}
	var ctlg string = ""
	for n_it := 0; n_it < news_it; n_it++ {
		CreateDir(savePath+news[n_it].Prefix+"/"+news[n_it].Subject+"/")
		saveNews(news[n_it], savePath+news[n_it].Prefix+"/"+news[n_it].Subject)
		ctlg += savePath+news[n_it].Prefix+"/"+news[n_it].Subject+"/"+news[n_it].NewId + ".sns"
		ctlg += "\n"
	}
	AppendFile(CATALOG_PATH+specialCoverage+"/", "path.snp", ctlg)
}
//获取新闻的url，直接保存到对应专题的path
func GetNewsUrls(Url string, specialCoverage string, savePath string) {
	rawHtml := GetHtml(Url)
	subjectName := GetSubjectName(rawHtml)
	if subjectName == "" {
		log.Errorln("No such subject. Ignore...")
		return
	}
	log.Infoln(subjectName)

	ret := newsRegExp.FindAllString(rawHtml, -1)
	subjectName = strings.Replace(subjectName, " ", "_", -1)
	news, news_it := GewNews(ret, subjectName, specialCoverage)

	log.Infoln("Special Coverage : ", specialCoverage)
	log.Infoln("Total : ", len(ret))
	log.Infoln("Saving news. Path : ", savePath)

	SaveNews(news, news_it, savePath, specialCoverage)
}

// 保存新闻，目录为对应专题，保存newid.sns
func saveNews(new New, savePath string) {
	var fileName string = new.NewId + ".sns"
	var buf string = new.Subject + "\n" + new.Title + "\n" + new.Time + "\n" + new.Content + "\n"
	SaveFile(savePath+"/", fileName, buf)
}

//页面样式1
func getNewsProcess1(Url string, Subject string, specialCoverage string) New {
	var new = New{"","","","","",""}
	doc, err := goquery.NewDocument(Url)
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
	if new.Prefix != "" && new.Subject != "" && new.Content != "" && new.Time != "" {
		return new
	}
	return New{SIGNAL,"","","","",""}
}

//页面样式2
func getNewsProcess2(Url string, Subject string, specialCoverage string) New {
	var new = New{"","","","","",""}
	doc, err := goquery.NewDocument(Url)
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
	new.Title = doc.Find(".wrap .part_01 .p_left #Article #artibodyTitle h1").First().Text()
	new.Time = doc.Find(".wrap .part_01 .p_left #Article #artibodyTitle span").First().Text()
	contentSelection := doc.Find(".wrap .part_01 .p_left #Article #artibody")
	for s_it := 0; s_it < contentSelection.Size(); s_it++ {
		content := contentSelection.Get(s_it)
		new.Content += goquery.NewDocumentFromNode(content).Find("p").Text()
	}
	if new.Prefix != "" && new.Subject != "" && new.Content != "" && new.Time != "" {
		return new
	}
	return New{SIGNAL,"","","","",""}
}

// TODO: query获取新闻时间、标题、内容
func getNews(Url string, Subject string, specialCoverage string) New {
	var new = New{"","","","","",""}
	new = getNewsProcess1(Url, Subject, specialCoverage)
	if new.Prefix != SIGNAL {
		return new
	}
	new = getNewsProcess2(Url, Subject, specialCoverage)
	if new.Prefix != SIGNAL {
		return new
	}
	return New{SIGNAL,"","","","",""}
}