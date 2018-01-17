package main

import (
    "github.com/go-martini/martini"
    "net/http"
    "io/ioutil"
    "strconv"
    "fmt"
)

const SERVER_ADDR = "192.168.216.201:11211"

func responseJSON(result bool, message interface{}) string {
    return fmt.Sprintf("{'result':%t, 'message':'%v'}", result, message)
}

func Add(params martini.Params, req *http.Request) string {

    key := params["key"]

    flag , err := strconv.Atoi(params["flag"])
    if err != nil{
        return responseJSON(false, "param flag error")
    }

    expire , err := strconv.Atoi(params["expire"])
    if err != nil{
        return responseJSON(false, "param expire error")
    }

    content , err := ioutil.ReadAll(req.Body)
    if err != nil{
        return responseJSON(false, "post content error")
    }

    if len(content) == 0{
        return responseJSON(false, "post content is empty")
    }

    contentString := string(content)
    client := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }
    result := client.Add(key, flag, expire, contentString)
    return responseJSON(result, "")
}

func Get(params martini.Params, req *http.Request) string {
    key := params["key"]
    client := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    response, result:= client.Get(key)
    if !result {
        return responseJSON(false, "get key error")
    }
    return responseJSON(true, response)
}

func Delete(params martini.Params, req *http.Request) string {
    key := params["key"]
    client  := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    result := client.Delete(key)
    return responseJSON(result, "")
}

func FlushAll(params martini.Params, req *http.Request) string {
    client  := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }
    result := client.FlushAll()
    return responseJSON(result, "")
}
