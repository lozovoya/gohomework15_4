package main

import (
	"github.com/lozovoya/gohomework15_3/pkg/middleware/logger"
	pages2 "github.com/lozovoya/gohomework15_4/pkg/pages"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/lozovoya/gohomework15_3/pkg/remux"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute (addr string) error {
	rmux := remux.New()
	loggerMd := logger.Logger
	pages := pages2.NewService()

	rmux.RegisterPlain(remux.GET, "/ok", http.HandlerFunc(pages.Ok), loggerMd)
	rmux.RegisterPlain(remux.PUT, "/AddPage", http.HandlerFunc(pages.AddPage), loggerMd)
	rmux.RegisterPlain(remux.GET, "/GetPages", http.HandlerFunc(pages.GetPages), loggerMd)

	log.Fatal(http.ListenAndServe(addr, rmux))

	return nil
}