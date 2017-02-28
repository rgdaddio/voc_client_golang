package main

import (
    "fmt"
    "os"
    "bytes"
    "database/sql"
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


    send_req("https://" + os.Args[1] + "/Anaina/v0/Register",  ja_buffer)

    send_req("https://" + os.Args[1] + "/Anaina/v0/HelloVoC",  ja_buffer)

    db, _ := sql.Open("sqlite3", "./foo.db")
    create_tables(db);
    //fmt.Printf("%s", db)


}
