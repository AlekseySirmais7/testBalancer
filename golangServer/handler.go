package main

import (
	"net/http"
)

const (
	workCoefficient = 0
)

func GetSimplePage(w http.ResponseWriter, r *http.Request, returnPage *string) {

	// check err case
	keys, ok := r.URL.Query()["param"]
	if ok && len(keys[0]) > 1 && keys[0] == "err" {
		w.WriteHeader(http.StatusInternalServerError)
		totalBadReqCount.Add(1)
		return
	}

	// some work emulation
	workTimes := 10000000 * workCoefficient
	result := 10
	for i := 0; i < workTimes; i++ {
		result = (result*i + 15) % 33
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(*returnPage))
	totalOkReqCount.Add(1)
	oneHandlerCountOk.WithLabelValues("200", r.RequestURI).Inc()
}
