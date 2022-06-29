package main

import (
	"fmt"
	"net/http"
)

func main (){
	http.HandleFunc("/hello", HelloHandleFunc)
	http.HandleFunc("/about", AboutHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func HelloHandleFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
}

func AboutHandleFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"r.URL.Path:, %s", r.URL.Path)
}