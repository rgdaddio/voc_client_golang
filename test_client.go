package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "crypto/tls"
    "bytes"
//    "encoding/json"
)

func build_header(){
     //return head
}

func build_json(schema string, tenant string, pubkey string) string{
     var buf bytes.Buffer
     buf.WriteString("{")
     buf.WriteString("\"ServerState\"")
     buf.WriteString(":")
     buf.WriteString("{")
     buf.WriteString("\"schemaName\"")
     buf.WriteString(":")
     buf.WriteString("\"")
     fmt.Println(buf.String())
     return buf.String()
}

func main() {
    
    if len(os.Args) != 5 {
       fmt.Fprintf(os.Stderr, "Usage: client <server> <schemaName> <tenantId> <publicKey>\n")
       os.Exit(1)
    } 
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    }

    client := &http.Client{Transport: tr}
    var url_buffer bytes.Buffer
    url_buffer.WriteString("https://")
    url_buffer.WriteString(os.Args[1])
    url_buffer.WriteString("/Anaina/v0/Register")
    jsnn := build_json(os.Args[1], os.Args[2], os.Args[3])
    fmt.Println(jsnn)
    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast	."}`)
    req, err := client.Post(url_buffer.String(), "application/json", bytes.NewBuffer(jsonStr))	
    //response, err := client.Get(url_buffer.String())
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer req.Body.Close()
        contents, err := ioutil.ReadAll(req.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}