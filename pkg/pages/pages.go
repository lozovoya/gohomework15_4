package pages

import (
	"encoding/json"
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
	//pages := make([]*pages2.PageDTO, 0)
	for _, page := range p.Pages {
		log.Println(page.Name)
		//pages.
	}
	w.Write([]byte("get pages"))
	//body, err :=  json.Marshal()
}


