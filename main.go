package main

import (
    "github.com/go-martini/martini"
    l4g "github.com/alecthomas/log4go"
)

func main() {

    l4g.LoadConfiguration("log4go.xml")
    l4g.Info("Application Start")
    defer l4g.Close()

    m := martini.Classic()
    m.Post("/action/add/:key/:flag/:expire", Add)
    m.Post("/action/set/:key/:flag/:expire", Set)
    m.Post("/action/replace/:key/:flag/:expire", Replace)
    m.Post("/action/get/:key", Get)
    m.Post("/action/delete/:key", Delete)
    m.Post("/action/delete/:key/:delay", Delete)
    m.Post("/action/flushall", FlushAll)
    m.Post("/action/incr/:key/:num", Incr)
    m.Post("/action/decr/:key/:num", Decr)

    m.Run()
}
