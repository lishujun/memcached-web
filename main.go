package main

import (
    "github.com/go-martini/martini"
)

func main() {
    m := martini.Classic()
    m.Post("/action/add/:key/:flag/:expire", Add)
    m.Post("/action/get/:key", Get)
    m.Post("/action/delete/:key", Delete)
    m.Post("/action/flushall", FlushAll)
    m.Run()
}
