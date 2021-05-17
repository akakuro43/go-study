package main

import (
    "fmt"
    "time"
    // "log"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
)

// var db nil

func main() {
    db, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("DB接続成功")
    }

    defer db.Close()

    // レコード追加
    error := db.Create(&Users{
        ID:       0004,
        Name:     "テスト太郎2",
        Age:      18,
        Address:  "東京都千代田区",
        UpdateAt: getDate(),
    }).Error
    if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println("データ追加成功")
    }
}

func getDate() string {
    const layout = "2006-01-02 15:04:05"
    now := time.Now()
    return now.Format(layout)
}
 
// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "localuser"
    PASS := "localpass"
    PROTOCOL := "tcp(db:3306)"
    DBNAME := "localdb"
 
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type Users struct {
    ID       int
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Address  string `json:"address"`
    UpdateAt string `json:"updateAt" sql:"not null;type:date"`
}