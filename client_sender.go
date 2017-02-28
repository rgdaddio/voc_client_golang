package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "crypto/tls"
    "bytes"
)
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