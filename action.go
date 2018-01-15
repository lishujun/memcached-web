package main

import (
    "github.com/go-martini/martini"
    //"io/ioutil"
    "net/http"
    //"strings"
    "io/ioutil"
    "strconv"
)

const SERVER_ADDR = "192.168.216.201:11211"

func Add(params martini.Params, req *http.Request) string {

    key := params["key"]
    flag , _ := strconv.Atoi(params["flag"])
    expire , _ := strconv.Atoi(params["expire"])

    content , err := ioutil.ReadAll(req.Body)
    if err != nil{
        return "read content error"
    }

    contentString := string(content)

    client := MakeClient(SERVER_ADDR)
    return client.Add(key, flag, expire, contentString)
}

func Get(params martini.Params, req *http.Request) string {
    key := params["key"]
    client := MakeClient(SERVER_ADDR)
    return client.Get(key)
}

func Delete(params martini.Params, req *http.Request) string {
    key := params["key"]
    client := MakeClient(SERVER_ADDR)
    return client.Delete(key)
}

func FlushAll(params martini.Params, req *http.Request) string {
    client := MakeClient(SERVER_ADDR)
    return client.FlushAll()
}
