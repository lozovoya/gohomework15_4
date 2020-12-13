package pages

import (
	"encoding/json"
	"errors"
	"github.com/lozovoya/gohomework15_3/pkg/remux"
	"github.com/lozovoya/gohomework15_4/pkg/pages/DTO"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	HttpReplyError = errors.New("http reply error")
)

type Page struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Pic     string    `json:"pic"`
	Article string    `json:"article"`
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

func (p *Service) SendReply(body interface{}, httpCode int, ContentType string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", ContentType)
	w.WriteHeader(httpCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return HttpReplyError
	}
	return nil
}

func (p *Service) AddPage(w http.ResponseWriter, r *http.Request) {

	var page *Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		log.Println(err)
		return
	}

	if (page.Name == "") || (page.Pic == "") || (page.Article == "") {
		err = p.SendReply("some field is empty", 400, "text/plain", w)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	if len(p.Pages) == 0 {
		page.Id = 1
	} else {
		page.Id = p.Pages[len(p.Pages)-1].Id + 1
	}

	page.Created = time.Now()
	p.Pages = append(p.Pages, page)
	err = p.SendReply(page, 201, "application/json", w)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (p *Service) GetPages(w http.ResponseWriter, r *http.Request) {
	pages := make([]dto.PagesDTO, len(p.Pages))
	if len(pages) == 0 {
		log.Println("no pages available")
		err := p.SendReply("no pages available", 200, "text/plain", w)
		if err != nil {
			log.Println(err)
			return
		}
	}

	for i, page := range p.Pages {
		pages[i].Id = page.Id
		pages[i].Name = page.Name
		pages[i].Pic = page.Pic
		pages[i].Created = page.Created
	}

	err := p.SendReply(pages, 200, "application/json", w)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (p *Service) GetPageById(w http.ResponseWriter, r *http.Request) {
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

			err := p.SendReply(respPage, 200, "application/json", w)
			if err != nil {
				log.Println(err)
				return
			}

			return
		}
	}

	err = p.SendReply("No page with such id", 200, "text/plain", w)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (p *Service) UpdatePageById(w http.ResponseWriter, r *http.Request) {
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

	var inPage *Page
	err = json.NewDecoder(r.Body).Decode(&inPage)
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

			err = p.SendReply(respPage, 200, "application/json", w)
			if err != nil {
				log.Println(err)
				return
			}

			return
		}
	}

	err = p.SendReply("No page with such id", 200, "plain/text", w)
	if err != nil {
		log.Println(err)
		return
	}

	return

}

func (p *Service) DeletePageById(w http.ResponseWriter, r *http.Request) {
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

	for i, singlePage := range p.Pages {
		if singlePage.Id == id {
			p.Pages = append(p.Pages[:i], p.Pages[i+1:]...)
			err = p.SendReply("page deleted", 204, "plain/text", w)
			if err != nil {
				log.Println(err)
				return
			}

			return
		}
	}

	err = p.SendReply("No page with such id", 200, "plain/text", w)
	if err != nil {
		log.Println(err)
		return
	}

	return

}
