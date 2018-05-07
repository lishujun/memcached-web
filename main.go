package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/sessions"
    l4g "github.com/alecthomas/log4go"
)

func main() {

    l4g.LoadConfiguration("log4go.xml")
    l4g.Info("Application Start")
    defer l4g.Close()

    InitConfigReader()

    m := martini.Classic()
    store := sessions.NewCookieStore([]byte("secret123"))
    m.Use(sessions.Sessions("my_session", store))

    m.Post("/user/login", Login)
    m.Delete("/user/login", Logout)
    m.Post("/action/keys/:key/:flag/:expire", Add)
    m.Put("/action/keys/:key/:flag/:expire", Set)
    m.Get("/action/keys/:key", Get)
    m.Delete("/action/keys/:key", Delete)
    m.Delete("/action/keys", FlushAll)
    m.Put("/action/keys/:key/incr/:num", Incr)
    m.Put("/action/keys/:key/decr/:num", Decr)

    m.Run()
}
