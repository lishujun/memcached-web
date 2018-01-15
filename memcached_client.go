package main

import (
    "net"
    "fmt"
    //"strconv"
)

type MemcachedClient struct {
    conn net.Conn
}

func MakeClient(connString string) *MemcachedClient{
    var client *MemcachedClient
    client = &MemcachedClient{}
    client.init(connString)
    return client
}

func (this *MemcachedClient) init (connString string){
    conn, err := net.Dial("tcp", connString)
    if err != nil{
        fmt.Print(err)
        panic(err)
    }

    this.conn = conn
}

func (this *MemcachedClient) callApi(command string) string {
    this.conn.Write([]byte(command))

    var response []byte = make([]byte, 1024)
    n, _ := this.conn.Read(response)
    return string(response[:n])
}

func (this *MemcachedClient) callSaveApi(command string, data string) string {

    this.conn.Write([]byte(command))

    this.conn.Write([]byte(data + "\r\n"))

    var response []byte = make([]byte, 1024)
    n, _ := this.conn.Read(response)
    return string(response[:n])
}

func (this *MemcachedClient) Add(key string, flag int, expire int, data string) string {
    command := fmt.Sprintf("add %s %d %d %d \r\n", key, flag, expire, len(data))
    return this.callSaveApi(command, data)
}

func (this *MemcachedClient) Get(key string) string{
    command := fmt.Sprintf("get %s \r\n", key)
    return this.callApi(command)
}

func (this *MemcachedClient) Delete(key string) string{
    command := fmt.Sprintf("delete %s \r\n", key)
    return this.callApi(command)
}

func (this *MemcachedClient) FlushAll() string{
    command := fmt.Sprintf("flush_all \r\n")
    return this.callApi(command)
}

