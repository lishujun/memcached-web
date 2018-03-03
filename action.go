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

//const SERVER_ADDR = "192.168.216.201:11211"

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

func Add(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]

    flag , err := strconv.Atoi(params["flag"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param flag error")
    }

    expire , err := strconv.Atoi(params["expire"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param expire error")
    }

    content , err := ioutil.ReadAll(req.Body)
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "post content error")
    }

    if len(content) == 0{
        return http.StatusBadRequest, responseJSON(false, "post content is empty")
    }

    contentString := string(content)
    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.Add(key, flag, expire, contentString)
    return http.StatusOK, responseJSON(result, "")
}

func Replace(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]

    flag , err := strconv.Atoi(params["flag"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param flag error")
    }

    expire , err := strconv.Atoi(params["expire"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param expire error")
    }

    content , err := ioutil.ReadAll(req.Body)
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "post content error")
    }

    if len(content) == 0{
        return http.StatusBadRequest, responseJSON(false, "post content is empty")
    }

    contentString := string(content)
    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.Replace(key, flag, expire, contentString)
    return http.StatusOK, responseJSON(result, "")
}

func Set(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]

    flag , err := strconv.Atoi(params["flag"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param flag error")
    }

    expire , err := strconv.Atoi(params["expire"])
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "param expire error")
    }

    content , err := ioutil.ReadAll(req.Body)
    if err != nil{
        return http.StatusBadRequest, responseJSON(false, "post content error")
    }

    if len(content) == 0{
        return http.StatusBadRequest, responseJSON(false, "post content is empty")
    }

    contentString := string(content)
    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.Set(key, flag, expire, contentString)
    return http.StatusOK, responseJSON(result, "")
}

func Get(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]
    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.Get(key)
    return http.StatusOK, responseJSON(true, result)
}

func Delete(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]
    //delay := params["delay"]

    client  := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.Delete(key)
    return http.StatusOK, responseJSON(result, "")
}

func FlushAll(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    client  := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    result := client.FlushAll()
    return http.StatusOK, responseJSON(result, "")
}

func Incr(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]
    num , err := strconv.Atoi(params["num"])

    if err != nil {
        return http.StatusBadRequest, responseJSON(false, "num error")
    }

    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    newNum, result := client.Incr(key, num)
    if ! result {
        return http.StatusInternalServerError, responseJSON(false, "incr error")
    }

    return http.StatusOK, responseJSON(true, newNum)
}

func Decr(params martini.Params, req *http.Request, session sessions.Session) (int,string) {

    if ! CheckAuth(session){
        return http.StatusUnauthorized, responseJSON(false, "No Auth")
    }

    key := params["key"]
    num , err := strconv.Atoi(params["num"])

    if err != nil {
        return http.StatusBadRequest, responseJSON(false, "num error")
    }

    client := MakeClient(ConfigReader.ConnString)
    if client == nil{
        return http.StatusInternalServerError, responseJSON(false, "make client error")
    }

    defer client.Close()
    newNum, result := client.Decr(key, num)
    if ! result {
        return http.StatusInternalServerError, responseJSON(false, "incr error")
    }

    return http.StatusOK, responseJSON(true, newNum)
}