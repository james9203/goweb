package main

import (
	"io"
	"net/http"
	"log"
)

var mux map[string]func(http.ResponseWriter,*http.Request)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &myHandler2{},
	}
	mux = make(map[string]func(http.ResponseWriter,*http.Request))
	mux["/hello"] = sayHello3
	mux["/bye"] = bye
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler2 struct{}

func (*myHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayHello3(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Hello world, this is version 3")
}

func bye(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"bye bye")
}
