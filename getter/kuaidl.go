package getter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/weaming/golib/fs"
	"github.com/weaming/proxy-crawler/storage"
)

func KDL() {
	pullURL := "http://www.kuaidaili.com/free/inha/"
	for i := 1; i < 50; i++ {
		currentPullURL := pullURL + strconv.Itoa(i)
		resp, err := http.Get(currentPullURL)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%v %v\n", currentPullURL, resp.Status)
		go analysisHTML(resp)
		time.Sleep(1 * time.Second)
	}
}

func analysisHTML(res *http.Response) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Println(err)
	}
	selection := doc.Find("#list tbody tr")
	selection.Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		port, err := strconv.Atoi(td.Nodes[1].FirstChild.Data)
		if err != nil {
			panic(err)
		}

		ipModel := storage.NewIP()
		ipModel.IP = td.Nodes[0].FirstChild.Data
		ipModel.Port = port
		ipModel.Protocol = td.Nodes[3].FirstChild.Data

		if storage.IsValidIP(ipModel) {
			ipModel.Usable = true
			saveToJson(ipModel)
		} else {
			log.Printf("%s:%v %v", ipModel.IP, ipModel.Port, false)
		}
	})
}

func saveToJson(ipModel *storage.IP) {
	b, err := json.Marshal(ipModel)
	fatalErr(err)
	text := string(b)
	fs.AppendToFile("crawler.txt", text)
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
