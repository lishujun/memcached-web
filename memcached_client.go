package main

import (
    "net"
    "fmt"

    l4g "github.com/alecthomas/log4go"
)

type MemcachedClient struct {
    conn net.Conn
}

func MakeClient(connString string) *MemcachedClient{

    conn, err := net.Dial("tcp", connString)
    if err != nil{
        l4g.Error("connection failed, reason: [%v]",  err)
        return nil
    }

    var client *MemcachedClient
    client = &MemcachedClient{}
    client.conn = conn
    return client
}

func (this *MemcachedClient) callApi(command string) string {

    _ , err := this.conn.Write([]byte(command))
    if err != nil{
        l4g.Error("send command '%s' failed, reason:[%v]", command, err)
        return ""
    }

    var response []byte = make([]byte, 1024 * 1024)
    n, err := this.conn.Read(response)
    if err != nil{
        l4g.Error("recv response failed, reason:[%v]", err)
        return ""
    }
    return string(response[:n])
}

func (this *MemcachedClient) callSaveApi(command string, data string) string {

    _ , err := this.conn.Write([]byte(command))
    if err != nil{
        l4g.Error("send command '%s' failed, reason: [%v]", command, err)
        return ""
    }

    _, err = this.conn.Write([]byte(data + "\r\n"))
    if err != nil{
        l4g.Error("send value failed, reason: [%v]", err)
        return ""
    }

    var response []byte = make([]byte, 1024)
    n, err := this.conn.Read(response)
    if err != nil{
        l4g.Error("recv response failed, reason: [%v]", err)
        return ""
    }
    return string(response[:n])
}

func (this *MemcachedClient) Add(key string, flag int, expire int, data string) bool {
    command := fmt.Sprintf("add %s %d %d %d \r\n", key, flag, expire, len(data))
    response  := this.callSaveApi(command, data)
    return response == "STORED"
}

func (this *MemcachedClient) Get(key string) (string, bool){
    command := fmt.Sprintf("get %s \r\n", key)
    response  :=  this.callApi(command)
    return response, true
}

func (this *MemcachedClient) Delete(key string) bool{
    command := fmt.Sprintf("delete %s \r\n", key)
    response := this.callApi(command)
    return response == "DELETED"
}

func (this *MemcachedClient) FlushAll() bool {
    command := fmt.Sprintf("flush_all \r\n")
    response := this.callApi(command)
    return response == "OK"
}

