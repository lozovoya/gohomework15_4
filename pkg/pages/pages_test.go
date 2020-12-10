package pages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lozovoya/gohomework15_3/pkg/remux"
	dto "github.com/lozovoya/gohomework15_4/pkg/pages/DTO"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)


func TestService_AddPage(t *testing.T) {

	rmux := remux.New()
	pages := NewService()

	rmux.RegisterPlain(remux.POST, "/pages", http.HandlerFunc(pages.AddPage))
	rmux.RegisterPlain(remux.GET, "/pages", http.HandlerFunc(pages.GetPages))

	type wants struct {
		page Page
		httpcode int
	}
	type args struct {
		method remux.Method
		path   string
		page dto.AddPageDTO
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{name: "PostPage", args: args{method: remux.POST, path: "/pages", page: dto.AddPageDTO{
			Name:    "Article1",
			Pic:     "http://www.url1.ru",
			Article: "text1text1text1",
		}}, want: wants{
			page: Page {
					Id:      1,
					Name:    "Article1",
					Pic:     "http://www.url1.ru",
					Article: "text1text1text1",
			},
			httpcode: 201,
		} },
		{name: "PostWrongPage", args: args{method: remux.POST, path: "/pages", page: dto.AddPageDTO{
			Name:    "",
			Pic:     "qwe",
			Article: "qwe",
		}}, want: wants{
			httpcode: 400,
			}	,
		},
	}
		for _, tt := range tests {
		page, _ := json.Marshal(tt.args.page)
		body := bytes.NewBuffer(page)
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, body)
		response := httptest.NewRecorder()
		rmux.ServeHTTP(response, request)

		if response.Code != tt.want.httpcode {
			t.Errorf("wrong responce code. got %v wanted %v", response.Code, tt.want.httpcode)
		}

		var got *Page
		json.Unmarshal(response.Body.Bytes(), &got)
		if (got.Id != tt.want.page.Id) || (got.Article != tt.want.page.Article) || (got.Pic != tt.want.page.Pic) || (got.Name != tt.want.page.Name){
			t.Error("got ", *got)
			t.Error("want ", tt.want.page)
		}
	}

}



func TestService_GetPages(t *testing.T) {

	rmux := remux.New()

	pages := NewService()
	var testPage = Page{
		Id:      1,
		Name:    "Article1",
		Pic:     "http://www.url1.ru",
		Article: "text1text1text1",
		Created: time.Now(),
	}
	pages.Pages = append(pages.Pages, &testPage)

	println(&testPage)
	fmt.Println(pages.Pages[0])

	rmux.RegisterPlain(remux.GET, "/pages", http.HandlerFunc(pages.GetPages))

	type wants struct {
		page dto.PagesDTO
		httpcode int
	}
	type args struct {
		method remux.Method
		path   string

	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{name: "GetAllPages", args: args{method: remux.GET, path: "/pages"}, want: wants{
			page: dto.PagesDTO {
				Id:      pages.Pages[0].Id,
				Name:    pages.Pages[0].Name,
				Pic:     pages.Pages[0].Pic,
				Created: pages.Pages[0].Created,
			},
			httpcode: 200,
		} },
	}
	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		rmux.ServeHTTP(response, request)

		if response.Code != tt.want.httpcode {
			t.Errorf("wrong responce code. got %v wanted %v", response.Code, tt.want.httpcode)
		}

		var got *dto.PagesDTO
		json.Unmarshal(response.Body.Bytes(), &got)
		if (got.Id != tt.want.page.Id) || (got.Created != tt.want.page.Created) || (got.Pic != tt.want.page.Pic) || (got.Name != tt.want.page.Name){
			t.Error("got ", *got)
			t.Error("want ", tt.want.page)
		}
	}

}


func TestService_GetPageById(t *testing.T) {

	rmux := remux.New()

	pages := NewService()
	var testPage = Page{
		Id:      1,
		Name:    "Article1",
		Pic:     "http://www.url1.ru",
		Article: "text1text1text1",
		Created: time.Now(),
	}
	pages.Pages = append(pages.Pages, &testPage)
	fmt.Println(pages.Pages[0].Created)

	regex, _ := regexp.Compile("^/pages/:(?P<id>\\d+)$")
	rmux.RegisterRegex(remux.GET, regex, http.HandlerFunc(pages.GetPageById))

	type wants struct {
		page dto.PageDTO
		httpcode int
	}
	type args struct {
		method remux.Method
		path   string

	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{name: "GetOnePage", args: args{method: remux.GET, path: "/pages/:1"}, want: wants{
			page: dto.PageDTO {
				Id:      pages.Pages[0].Id,
				Name:    pages.Pages[0].Name,
				Pic:     pages.Pages[0].Pic,
				Created: pages.Pages[0].Created,
			},
			httpcode: 200,
		} },
	}
	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		rmux.ServeHTTP(response, request)

		if response.Code != tt.want.httpcode {
			t.Errorf("wrong responce code. got %v wanted %v", response.Code, tt.want.httpcode)
		}

		var got *dto.PagesDTO
		json.Unmarshal(response.Body.Bytes(), &got)
		if (got.Id != tt.want.page.Id) || (got.Pic != tt.want.page.Pic) || (got.Name != tt.want.page.Name){
			t.Error("got ", *got)
			t.Error("want ", tt.want.page)
		}
	}

}