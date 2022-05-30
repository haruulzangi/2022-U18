package main

import (
	"log"
	"net/http"
)

func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Powered-By", "HZU18{M0ng0l1@n_f00d_v3ry_t@stY}")
		fs.ServeHTTP(w, r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", addHeaders(fs))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
