package main

import (
    "testing"
)

func TestExists(t *testing.T){
    InitSession(&Session)
    sessionid := Session.StartNewSession()
    exists := Session.IsLogin(sessionid)
    if exists != true{
        t.Error("Exists Error")
    }
}