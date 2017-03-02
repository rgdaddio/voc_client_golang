package main

import (
    "fmt"
    "os"
    "bytes"
    "bufio"
    "strings"
//    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

func build_header(){
     //return head
}

func main() {

    if len(os.Args) != 5 {
       fmt.Fprintf(os.Stderr, "Usage: client <server> <schemaName> <tenantId> <publicKey>\n")
       os.Exit(1)
    }

    var ja_buffer bytes.Buffer
    jsnn := build_json(os.Args[2], os.Args[3], os.Args[4])
    ja_buffer.WriteString(jsnn)

      reg_request("https://" + os.Args[1] + "/Anaina/v0/Register", ja_buffer)
      //fmt.Printf("%s", db)

    reader := bufio.NewReader(os.Stdin)
    for{
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if (text == "test") {
	   fmt.Printf("%s\n", text)
	   }else if (text == "hello") {
		 send_req("https://" + os.Args[1] + "/Anaina/v0/HelloVoC",  ja_buffer)
	   }else{
		fmt.Printf("Unknown command %s\n", text)
	}
      }
}
