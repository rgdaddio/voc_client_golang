package main

import (
    "fmt"
    "os"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

func create_tables(db *sql.DB){
  // VOC_USER TABLE
  stmt, _ := db.Prepare("create table if not exists voc_user" +
                          "(userid text, password text," +
                          "device_id text, platform text," +
                          "device_type text, access_token text," +
                          "refresh_token text, voc_id text," +
                          "congestion_detection text, ads_frequency text," +
                          "daily_quota integer, daily_manifest integer," +
                          "daily_download_wifi integer, daily_download_cellular integer," +
                          "congestion text, sdk_capabilities text," +
                          "max_content_duration integer, play_ads text," +
                          "skip_policy_first_time text, tod_policy text," +
                          "token_expiration integer, server text," +
                          "server_state text, my_row integer primary key autoincrement)")
  _, err := stmt.Exec()
  if err != nil { panic(err) }

  // PROVIDER TABLE
  stmt, _ = db.Prepare("create table if not exists provider " +
                        " (name text unique, contentprovider text, subscribed integer)" )
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //CATEGORY TABLE
  stmt, _ = db.Prepare("create table if not exists category (name text unique,subscribed integer)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //UUID TABLE
  stmt, _ = db.Prepare("create table if not exists uuid_table (uuid text)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //PLAYING TABLE
  stmt, _ = db.Prepare("create table if not exists playing (unique_Id text,timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //CONTENT_STATUS
  stmt, _ = db.Prepare("create table if not exists content_status (download_time text,download length integer,download_duration real,eviction_info text,user_rating int,unique_id text, my_row integer primary key autoincrement)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //CONSUMPTION STATUS
  stmt, _ = db.Prepare("create table if not exists consumption_status (watch_time int,watchstart integer,watchend int,my_row integer primary key autoincrement)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //AD CONSUMPTION
  stmt, _ = db.Prepare(" create table if not exists ad_consumption_status (adurl text,duration int, starttime integer,stopposition int, clicked int,unique_id text, my_row integer primary key autoincrement)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }

  //cache_manifest
  stmt, _ = db.Prepare(" create table if not exists cache_manifest " +
    "( local_file text, local_thumbnail text, " +
       " local_nfo text, video_size integer, " +
       " thumbnail_size integer, download_date integer, " +
       " content_provider text, category text, " +
       " unique_id text, summary text, " +
       " title text, duration integer, " +
       " timestamp integer, sdk_metadata text, " +
       " streams text,   ad_server_url text, " +
       " tags text, priority integer, " +
       " object_type text, thumb_attrs text, " +
       " object_attrs text, children text, " +
       " policy_name text, key_server_url text, " +
       " save integer default 0, my_row integer primary key autoincrement)")
  stmt.Exec()
  _, err = stmt.Exec()
  if err != nil { panic(err) }
  fmt.Fprintf(os.Stdout, "Database Table Creation Complete\n")
}