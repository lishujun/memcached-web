package main

import (
    "net/http"
    "io/ioutil"
    "strconv"
    "encoding/json"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/sessions"
    l4g "github.com/alecthomas/log4go"
)

const SERVER_ADDR = "192.168.216.201:11211"

func responseJSON(result bool, message interface{}) string {

    response := map [string]interface{}{
        "result" : result,
        "message" : message,
    }
    bytes , err := json.Marshal(response)
    if err == nil{
        msg := string(bytes)
        l4g.Info("output : %s ", msg)
        return msg
    }

    l4g.Error("convert to json error , reason [%v]" , err)
    return "{\"result\": false, \"message\":\"object to json error\"}"


}

func Add(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

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

func Replace(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

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
    result := client.Replace(key, flag, expire, contentString)
    return responseJSON(result, "")
}

func Set(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

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
    result := client.Set(key, flag, expire, contentString)
    return responseJSON(result, "")
}

func Get(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

    key := params["key"]
    client := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    result := client.Get(key)
    return responseJSON(true, result)
}

func Delete(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

    key := params["key"]
    //delay := params["delay"]

    client  := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    result := client.Delete(key)
    return responseJSON(result, "")
}

func FlushAll(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

    client  := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }
    result := client.FlushAll()
    return responseJSON(result, "")
}

func Incr(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

    key := params["key"]
    num , err := strconv.Atoi(params["num"])

    if err != nil {
        return responseJSON(false, "num error")
    }

    client := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    newNum, result := client.Incr(key, num)
    if ! result {
        return responseJSON(false, "incr error")
    }

    return responseJSON(true, newNum)
}

func Decr(params martini.Params, req *http.Request, session sessions.Session) string {

    if ! CheckAuth(session){
        return responseJSON(false, "没有权限")
    }

    key := params["key"]
    num , err := strconv.Atoi(params["num"])

    if err != nil {
        return responseJSON(false, "num error")
    }

    client := MakeClient(SERVER_ADDR)
    if client == nil{
        return responseJSON(false, "make client error")
    }

    newNum, result := client.Decr(key, num)
    if ! result {
        return responseJSON(false, "incr error")
    }

    return responseJSON(true, newNum)
}