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

func send_req(url string, data bytes.Buffer){
  fmt.Printf("%s\n", url)
  tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
  }
  client := &http.Client{Transport: tr}
  req, err := client.Post(url, "application/json", &data)

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
      fmt.Fprintf(os.Stdout, "Response\n")
      fmt.Printf("%s\n", string(contents))
  }
}


func main() {

    if len(os.Args) != 5 {
       fmt.Fprintf(os.Stderr, "Usage: client <server> <schemaName> <tenantId> <publicKey>\n")
       os.Exit(1)
    }

    var ja_buffer bytes.Buffer
    jsnn := build_json(os.Args[2], os.Args[3], os.Args[4])
    ja_buffer.WriteString(jsnn)
    send_req("https://" + os.Args[1] + "/Anaina/v0/Register",  ja_buffer)

    send_req("https://" + os.Args[1] + "/Anaina/v0/HelloVoC",  ja_buffer)

}
