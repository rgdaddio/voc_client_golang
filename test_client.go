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

    reg_request(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

    //fmt.Printf("%s", db)

    var ja_buffer bytes.Buffer
    reader := bufio.NewReader(os.Stdin)
    for{
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if (text == "test") {
	   fmt.Printf("%s\n", text)
	   }else if (text == "hello") {
		 send_req("https://" + os.Args[1] + "/Anaina/v0/HelloVoC",  ja_buffer)
	   }else if (text == "status") {
		 send_req("https://" + os.Args[1] + "/Anaina/v0/Status",  ja_buffer)
	   }else{
		fmt.Printf("Unknown command %s\n", text)
	}
      }
}
