package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Hello Sir</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Hello server!!!")

	http.ListenAndServe(":8080",nil)
}