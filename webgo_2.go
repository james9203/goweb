package main

import (
	"net/http"
	"io"
	"log"
)

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/",&myHandler{})
	mux.Handle("/hello",&sayHello2{})
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
