package main

import (
    "net"
    "fmt"
    "strconv"

    l4g "github.com/alecthomas/log4go"
    "strings"
)

type CachedResult struct {
    Key string
    Type int
    Expire int
    Length int
    Data interface{}
}

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

    l4g.Info("send command '%s'", command)

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
    l4g.Info("recv response '%s'", response[:n])
    return string(response[:n])
}

func (this *MemcachedClient) callSaveApi(command string, data string) string {

    l4g.Info("send command '%s'", command)
    l4g.Debug("save data '%s'", data)
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
    l4g.Info("recv response '%s'", response[:n])
    return string(response[:n])
}

func (this *MemcachedClient) parse(response string) []CachedResult {

    var resultList []CachedResult = []CachedResult{}
    slice := response[:]
    for{
        i := strings.IndexAny(slice, "\r\n")
        if i != -1{
            if len(slice[:i]) > 5 && slice[:i][:5] == "VALUE"{
                var result CachedResult
                header := strings.Split(slice[:i], " ")
                result.Key = header[1]
                result.Type, _ = strconv.Atoi(header[2])
                result.Length, _ = strconv.Atoi(header[3])
                result.Expire = -1
                result.Data = slice[i+2:][:result.Length]
                resultList = append(resultList, result)

                slice = slice[i + 2 + result.Length + 2:]
                continue
            }
        }

        break
    }
    return resultList
}

func (this *MemcachedClient) Close(){
    if(this.conn != nil){
        l4g.Debug("close conn")
        this.conn.Close()
    }
}

func (this *MemcachedClient) Add(key string, flag int, expire int, data string) bool {
    command := fmt.Sprintf("add %s %d %d %d \r\n", key, flag, expire, len(data))
    response  := this.callSaveApi(command, data)
    return response == "STORED\r\n"
}

func (this *MemcachedClient) Replace(key string, flag int, expire int, data string) bool {
    command := fmt.Sprintf("replace %s %d %d %d \r\n", key, flag, expire, len(data))
    response  := this.callSaveApi(command, data)
    return response == "STORED\r\n"
}

func (this *MemcachedClient) Set(key string, flag int, expire int, data string) bool {
    command := fmt.Sprintf("set %s %d %d %d \r\n", key, flag, expire, len(data))
    response  := this.callSaveApi(command, data)
    return response == "STORED\r\n"
}

func (this *MemcachedClient) Get(key string) []CachedResult{
    command := fmt.Sprintf("get %s \r\n", key)
    response  :=  this.callApi(command)
    return this.parse(response)
}

func (this *MemcachedClient) Delete(key string) bool{
    command := fmt.Sprintf("delete %s \r\n", key)
    response := this.callApi(command)
    return response == "DELETED\r\n"
}

func (this *MemcachedClient) FlushAll() bool {
    command := fmt.Sprintf("flush_all \r\n")
    response := this.callApi(command)
    return response == "OK\r\n"
}

func (this *MemcachedClient) Incr(key string, num int) (int, bool){
    command := fmt.Sprintf("incr %s %d\r\n", key, num)
    response := this.callApi(command)
    n, err := strconv.Atoi(response)
    if err != nil{
        return -1, false
    }
    return n, true
}

func (this *MemcachedClient)Decr(key string, num int) (int,bool) {
    command := fmt.Sprintf("decr %s %d\r\n", key, num)
    response := this.callApi(command)
    n, err := strconv.Atoi(response)
    if err != nil{
        return -1, false
    }
    return n, true
}

