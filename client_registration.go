package main

import  (
    "database/sql"
    "bytes"
    "encoding/json"
    "fmt"
)



func reg_request(server string, schema string, tenant string, public_key string) {
    db, _ := sql.Open("sqlite3", "./foo.db")

    var ret bool = validate_user_for_reg(db)
    if !ret {

      var ja_buffer bytes.Buffer
      jsnn := build_reg_json(schema, tenant, public_key)
      ja_buffer.WriteString(jsnn)
      data := ja_buffer
      req_str := "https://" + server + "/Anaina/v0/Register"

      create_tables(db);
      sret := send_req(req_str,  data)
      fmt.Printf("%t\n", ret);
      res := RegistrationResponse{}
      json.Unmarshal([]byte(sret), &res)
      fmt.Println(res)
      insert_voc_user(db, res)
    }else{
      println("\nUser is already registered\n")
    }
}
