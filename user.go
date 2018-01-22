package main

import (
    "net/http"
    "github.com/martini-contrib/sessions"
    l4g "github.com/alecthomas/log4go"
)


func Login(req *http.Request, session sessions.Session) string {

    username := req.PostFormValue("username")
    password := req.PostFormValue("password")
    if username == "" || password == ""{
        return responseJSON(false, "用户名密码不能为空")
    }

    if username != ConfigReader.Username || password != ConfigReader.Password{
        l4g.Debug(ConfigReader.Username + ":" + ConfigReader.Password)
        return responseJSON(false, "用户名密码不正确")
    }

    session.Set("user", username)
    l4g.Info("user '%s' at '%s' login", username, req.RemoteAddr)
    return responseJSON(true, "")
}

func Logout(req *http.Request, session sessions.Session) string {

    l4g.Info("user '%s' at '%s' logout", session.Get("user"), req.RemoteAddr)
    session.Delete("user")
    return responseJSON(true, "")
}

func CheckAuth(session sessions.Session) bool {
    return session.Get("user") != nil
}

//func CheckAuth(req *http.Request, session sessions.Session) string {
//    l4g.Info("user '%s' at '%s' check auth", session.Get("user"), req.RemoteAddr)
//    return  responseJSON(session.Get("user") != nil, "")
//}
