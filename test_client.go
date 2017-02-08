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
    var ja_buffer bytes.Buffer
    url_buffer.WriteString("https://")
    url_buffer.WriteString(os.Args[1])
    url_buffer.WriteString("/Anaina/v0/Register")
    jsnn := build_json(os.Args[2], os.Args[3], os.Args[4])
    ja_buffer.WriteString(jsnn)
    req, err := client.Post(url_buffer.String(), "application/json", &ja_buffer)	
    
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
