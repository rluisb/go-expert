package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	//We can use an implemented function or an anonymous function
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", blog{title: "myBlog"})
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
