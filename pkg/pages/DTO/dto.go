package pages

import "time"

type AddPageDTO struct {
	Name string `json:"name"`
	Pic string `json:"pic"`
	Article string `json:"article"`
}

type PageDTO struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Pic string `json:"pic"`
	Article string `json:"article"`
	Created time.Time `json:"created"`
}
