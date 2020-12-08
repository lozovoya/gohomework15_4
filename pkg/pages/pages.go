package pages

import (
	"encoding/json"
	"github.com/lozovoya/gohomework15_3/pkg/remux"
	"github.com/lozovoya/gohomework15_4/pkg/pages/DTO"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	Article string `json:"article"`
	Created time.Time `json:"created"`
}

type Service struct {
	Pages []*Page
}

func NewService() *Service {
	return &Service{
		Pages: nil,
	}
}



func (p *Service) SendReply (respBody []byte, httpCode int, ContentType string, w http.ResponseWriter,) {
	w.Header().Set("Content-Type", ContentType)
	w.WriteHeader(httpCode)
	_, err := w.Write(respBody)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (p *Service) Ok (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (p *Service) AddPage (w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var page *Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		log.Println(err)
		return
	}
	page.Id = len(p.Pages) + 1
	page.Created = time.Now()

	p.Pages = append(p.Pages, page)
	respBody, err := json.Marshal(page)
	if err != nil {
		log.Println(err)
		return
	}
	p.SendReply(respBody, 201, "application/json", w)
	return
}

func (p *Service) GetPages (w http.ResponseWriter, r *http.Request) {
	pages := make([]dto.PagesDTO, len(p.Pages))
	if len(pages) == 0 {
		log.Println("no pages available")
		p.SendReply([]byte("no pages available"), 200, "text/plain", w)
	}

	for i, page := range p.Pages {
		pages[i].Id = page.Id
		pages[i].Name = page.Name
		pages[i].Pic = page.Pic
		pages[i].Created = page.Created
	}

	respBody, err := json.Marshal(pages)
	if err != nil {
		log.Println(err)
		return
	}
	p.SendReply(respBody, 200, "application/json", w)
	return
}

func (p *Service) GetPageById (w http.ResponseWriter, r *http.Request) {
	params, err := remux.PathParams(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	id, err := strconv.Atoi(params.Named["id"])
	if err != nil {
		log.Println(err)
		return
	}

	for _, singlePage := range p.Pages {
		if singlePage.Id == id {
			var respPage dto.PageDTO
			respPage.Id = id
			respPage.Name = singlePage.Name
			respPage.Pic = singlePage.Pic
			respPage.Article = singlePage.Article
			respPage.Created = singlePage.Created

			respBody, err := json.Marshal(respPage)
			if err != nil {
				log.Println(err)
				return
			}

			p.SendReply(respBody, 200, "application/json", w)
			return
		}
	}

	respBody, err := json.Marshal("No page with such id")
	if err != nil {
		log.Println(err)
		return
	}
	p.SendReply(respBody, 200, "text/plain", w)
	return
}

func (p *Service) UpdatePageById (w http.ResponseWriter, r *http.Request) {
	params, err := remux.PathParams(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	id, err := strconv.Atoi(params.Named["id"])
	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var inPage *Page
	err = json.Unmarshal(body, &inPage)
	if err != nil {
		log.Println(err)
		return
	}

	for _, singlePage := range p.Pages {
		if singlePage.Id == id {

			singlePage.Name = inPage.Name
			singlePage.Pic = inPage.Pic
			singlePage.Article = inPage.Article

			var respPage dto.PageDTO
			respPage.Id = id
			respPage.Name = singlePage.Name
			respPage.Pic = singlePage.Pic
			respPage.Article = singlePage.Article
			respPage.Created = singlePage.Created

			respBody, err := json.Marshal(respPage)
			if err != nil {
				log.Println(err)
				return
			}

			p.SendReply(respBody, 200, "application/json", w)
			return
		}
	}

	respBody, err := json.Marshal("No page with such id")
	if err != nil {
		log.Println(err)
		return
	}
	p.SendReply(respBody, 200, "plain/text", w)
	return

}

