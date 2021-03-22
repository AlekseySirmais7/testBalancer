package main

import "github.com/prometheus/client_golang/prometheus"

var totalOkReqCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "req_ok_total",
	Help: "Count of total requests of this server",
})

var totalBadReqCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "req_bad_total",
	Help: "Count of total requests of this server",
})

var oneHandlerCountOk = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "oneHandlerCountOk",
}, []string{"status", "path"})

var oneHandlerCountBad = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "oneHandlerCountBad",
}, []string{"status", "path"})
