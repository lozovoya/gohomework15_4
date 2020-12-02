package pages

import (
	"encoding/json"
	"github.com/lozovoya/gohomework15_4/pkg/pages/DTO"
	"io/ioutil"
	"log"
	"net/http"
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
	w.WriteHeader(200)
	w.Write([]byte("page is added"))
	return
}

func (p *Service) GetPages (w http.ResponseWriter, r *http.Request) {
	pages := make([]dto.PagesDTO, len(p.Pages))
	if len(pages) == 0 {
		log.Println("no pages available")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("no pages available"))
	}

	for i, page := range p.Pages {
		log.Println(page.Name)
		pages[i].Id = page.Id
		pages[i].Name = page.Name
		pages[i].Pic = page.Pic
		pages[i].Created = page.Created
	}

	body, err := json.Marshal(pages)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

	return
}


