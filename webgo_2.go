package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/",&myHandler{})
	mux.Handle("/hello",&sayHello2{})
	wd, err_wd := os.Getwd()
	if err_wd != nil{
		log.Fatal(err_wd)
	}
	mux.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir(wd))))
	err := http.ListenAndServe(":8080",mux)
	if err != nil{
		log.Fatal(err)
	}
}

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter ,r *http.Request)  {
	io.WriteString(w,"URL:"+r.URL.String())
}

type sayHello2 struct {

}
func (*sayHello2)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Hello world, this is version 2")
}
