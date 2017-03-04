package main

import (
    "fmt"
    "os"
    "bytes"
    "bufio"
    "strings"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

func build_header(){
     //return head
}

func status_req(url string, schema string, tenant string, db *sql.DB){
    voc_info := get_voc_info(db)
    var ja_buffer bytes.Buffer
    jsnn := build_status_json(schema, tenant, voc_info)
    ja_buffer.WriteString(jsnn)
    data := ja_buffer
    ret := send_req(url,  data)
    fmt.Printf("%t\n", ret);
}

func manifest_req(url string, schema string, tenant string, db *sql.DB){
     voc_info := get_voc_info(db)
     var ja_buffer bytes.Buffer
     jsnn := build_manifest_json(schema, tenant, voc_info)
     ja_buffer.WriteString(jsnn)
     data := ja_buffer
     ret := send_req(url, data)
     fmt.Printf("%t\n", ret)
}

func main() {

    db, _ := sql.Open("sqlite3", "./foo.db")

    if len(os.Args) != 5 {
       fmt.Fprintf(os.Stderr, "Usage: client <server> <schemaName> <tenantId> <publicKey>\n")
       os.Exit(1)
    }

    reg_request(os.Args[1], os.Args[2], os.Args[3], os.Args[4], db)

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
		 status_req("https://" + os.Args[1] + "/Anaina/v0/Status", os.Args[2], os.Args[3], db)
           }else if (text == "cacheFill") {
	   	 manifest_req("https://" + os.Args[1] + "/Anaina/v0/Download-Manifest", os.Args[2], os.Args[3], db)	   
	   }else{
		fmt.Printf("Unknown command %s\n", text)
	}
      }
}
