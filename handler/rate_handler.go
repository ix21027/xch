package handler

import (
	"encoding/json"
	"net/http"
	"xch/service"
)

func getRate(w http.ResponseWriter, r *http.Request) {
	rate, err := service.GetRate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}
