package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World! In GOLang"))
	})

	_ = http.ListenAndServe("0.0.0.0:8080", nil)
}
