package main

import  (
    "database/sql"
    "bytes"
    "encoding/json"
    "fmt"
)



func reg_request(req_str string, data bytes.Buffer) {
    db, _ := sql.Open("sqlite3", "./foo.db")

    var ret bool = validate_user_for_reg(db)
    if !ret {
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
