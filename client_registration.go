package main

import  (
    "database/sql"
    "bytes"
    "encoding/json"
    "fmt"
)



func reg_request(req_str string, data bytes.Buffer) {
    db, _ := sql.Open("sqlite3", "./foo.db")
    create_tables(db);
    ret := send_req(req_str,  data)
    println("%s\n", ret);
    res := RegistrationResponse{}
    json.Unmarshal([]byte(ret), &res)
    fmt.Println(res)
    insert_voc_user(db, res)
}
