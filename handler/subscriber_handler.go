package handler

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"xch/db"

	"github.com/mattn/go-sqlite3"
)

func addSubscriber(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if _, err := mail.ParseAddress(email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err := db.AddSubscriber(email)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint{
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(struct{ Email string }{email})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
