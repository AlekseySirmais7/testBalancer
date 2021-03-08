package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const (
	goServerPortStr = ":8080"
	serverNumberStr = "1"
	pagePath        = "./WikiGoPage.html"
)

func main() {

	returnPageHtml, errRead := readFile(pagePath)
	if errRead != nil {
		log.Fatal(errRead)
	}

	prometheus.MustRegister(totalOkReqCount, oneHandlerCountOk)
	prometheus.MustRegister(totalBadReqCount, oneHandlerCountBad)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		GetSimplePage(w, r, &returnPageHtml)
	})

	log.Println("Server №" + serverNumberStr)
	log.Println("Start serving " + goServerPortStr)
	log.Fatal(http.ListenAndServe(goServerPortStr, nil))
}
