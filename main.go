package main

import (
    "github.com/go-martini/martini"
)

func main() {
    m := martini.Classic()
    m.Get("/acton/Add/:key", Add)
    m.Get("/action/Get/:key", Get)
    m.Run()
}