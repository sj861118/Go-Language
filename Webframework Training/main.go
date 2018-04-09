package main

import(
  "fmt"
)



func main(){
  s := NewServer()

  s.HandleFunc("GET","/", func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "welcome")
  })

  s.Run(":9000")
/*
  r := &router{make(map[string]map[string]HandlerFunc)}


  r.HandleFunc("GET","/",logHandler(staticHandler(func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "welcome")
  })))

  r.HandleFunc("GET","/about",logHandler(func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "about")
  }))

  r.HandleFunc("GET","/users/:id",logHandler(recoverHandler(func(c *Context){
    if c.Params["id"] == "0" {
      panic("id is zero")
    }
    fmt.Fprintln(c.ResponseWriter, "retrieve user",c.Params["id"])
  })))

  r.HandleFunc("GET","/users/:user_id/addresses/:address_id",logHandler(func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "retrieve user's address")
  }))

  r.HandleFunc("POST","/users",logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "create user")
  })))))


  r.HandleFunc("POST","/users/:user_id/addresses",logHandler(func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "create address")
  }))


  //8080 포트로 웹서버 구동
  http.ListenAndServe(":9000",r)
  */
}
