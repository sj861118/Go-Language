package main

import(
  "fmt"
  "time"
)

type User struct{
  Id string
  AddressId string
}

func main(){
  s := NewServer()

  s.HandleFunc("GET","/", func(c *Context){
    fmt.Fprintln(c.ResponseWriter, "welcome")
  })

  s.HandleFunc("GET", "/users/:id", func(c *Context){
    u := User{Id:c.Params["id"].(string)}
    c.RenderXml(u)
  })

  s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context){
    u := User{Id:c.Params["user_id"].(string), AddressId:c.Params["address_id"].(string)}
    c.RenderJson(u)
  })

  s.HandleFunc("GET", "/", func(c *Context){
    c.RenderTemplate("/public/index.html", map[string]interface{}{"time": time.Now()})
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
