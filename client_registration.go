package main

import  (
    "database/sql"
    "bytes"
)

func reg_request(req_str string, data bytes.Buffer) {
    db, _ := sql.Open("sqlite3", "./foo.db")
    create_tables(db);
    ret := send_req(req_str,  data)
    println("%s\n", ret);
}