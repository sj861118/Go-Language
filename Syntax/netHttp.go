package main

import (
  "fmt"
  "net/http"
)

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "welcome")
  })

  http.ListenAndServe(":9000",nil)
}
