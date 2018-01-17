package main

import (
    "github.com/go-martini/martini"
    l4g "github.com/alecthomas/log4go"
)

func main() {

    l4g.LoadConfiguration("configuration.xml")
    l4g.Info("Application Start")
    defer l4g.Close()

    m := martini.Classic()
    m.Post("/action/add/:key/:flag/:expire", Add)
    m.Post("/action/get/:key", Get)
    m.Post("/action/delete/:key", Delete)
    m.Post("/action/flushall", FlushAll)
    m.Run()
}
