package main

import (
    "github.com/go-martini/martini"
    //"io/ioutil"
    "net/http"
    //"strings"
)

func Add(params martini.Params, req *http.Request) string {
    return "add"
}

func Get(params martini.Params, req *http.Request) string {
    key := params["key"]

    command := "get " + key + "\r\n"
    client := MakeClient("127.0.0.1 11211")
    return client.callApi(command)
}
