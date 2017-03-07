package main

import (
    "fmt"
    "database/sql"
    "encoding/json"
    _"github.com/mattn/go-sqlite3"
)


func handle_manifest(jstr string, db *sql.DB){
     var datas []ContentManifest
     json.Unmarshal([]byte(jstr), &datas)
     fmt.Printf("# of content manifests: %d\n ",len(datas))
      for i := range datas {
        fmt.Println(datas[i].Title)
        fmt.Println(datas[i].Streams)
        insert_content_manifest(db, datas[i])
      }
}