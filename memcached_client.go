package main

import (
    "net"
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

func (this *MemcachedClient) callSaveApi(
    command string, data string) string {

    this.conn.Write([]byte(command))

    this.conn.Write([]byte(data))

    var response []byte = make([]byte, 1024)
    n, _ := this.conn.Read(response)
    return string(response[:n])
}

