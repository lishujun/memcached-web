package main

import (
    "testing"
)

func TestExists(t *testing.T){
    InitSessionCollection(&Session)
    sessionid := Session.StartNewSession()
    exists := Session.IsLogin(sessionid)
    if exists != true{
        t.Error("Exists Error")
    }
}