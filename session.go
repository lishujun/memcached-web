package main

import (
    "time"
    "fmt"

    "github.com/satori/go.uuid"
    l4g "github.com/alecthomas/log4go"
)

type SESSION struct{
    sessions map[string] time.Time
}

var Session SESSION

func InitSessionCollection(session *SESSION) {
    l4g.Info("Init Session Collection")
    *session = SESSION{}
    (*session).sessions = make(map[string]time.Time)
}

func (session *SESSION) StartNewSession () string {

    v , _ := uuid.NewV4()
    sessionId := fmt.Sprintf("%s", v)
    session.sessions[sessionId] = time.Now()
    return sessionId
}

func (session *SESSION) IsLogin (sessionId string) bool {
    lastDate, ok := session.sessions[sessionId]
    if !ok{
        return false
    }

    if time.Now().Sub(lastDate).Minutes() > 20{
        delete(session.sessions, sessionId)
        return false
    }

    session.sessions[sessionId] = time.Now()
    return true
}



