package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "io"
    "os"
    "crypto/tls"
    "bytes"
)
func send_req(url string, data bytes.Buffer) string{
  fmt.Printf("%s\n", url)
  var ret string = "error"
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
      ret = string(contents)
  }
  return ret
}

/**
Download at url to disk with given file name
**/
func download(download_url string, file_name string){
  //create an instance of the file, if error it probably already exists
  //in the future do we clear the file?
  out, err := os.Create(file_name)
  if err != nil {
    fmt.Println("Error trying to open ", file_name, " error:", err)
    return
  }

  //defer closing the file until the function returns
  defer out.Close()

  //do a Get of Download url 
  // For now I think we assume normal get works without building TLSclient stuff?
  resp, err := http.Get(download_url)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer resp.Body.Close()
  n, err := io.Copy(out, resp.Body)

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("donwload %s complete %d bytes\n", file_name, n)
}
