package handlers

import (
	"io/ioutil"
	"learngo/fivem/episode1/storage"
	"net/http"
)

//PutKey handler
func PutKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading put body", http.StatusBadRequest)
			return
		}
		if err := db.Set(key, val); err != nil {
			http.Error(w, "error set", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
