// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo/controller"
	"todo/model"

	"github.com/joho/godotenv"
)

func main() {

	dbUser, dbPassword, dbHost, dbPort, dbName := loadDB()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// データベースとの接続を確立する
	err := model.InitDB(dbURI)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/todo", controller.GetTodos)
	http.HandleFunc("/todo/add", controller.AddTodo)
	http.HandleFunc("/todo/update", controller.UpdateTodo)
	http.HandleFunc("/todo/delete", controller.DeleteTodo)

	http.ListenAndServe(":8080", nil)
}

func loadDB() (string, string, string, string, string) {
	// .env ファイルの読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数からデータベース接続情報を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return dbUser, dbPassword, dbHost, dbPort, dbName
}

// DB接続を以下のように記載する方法もある。
// jst, err := time.LoadLocation("Asia/Tokyo")
// if err != nil {
//     // エラーハンドリング
// }
// c := mysql.Config{
//     DBName:    "db",
//     User:      "user",
//     Passwd:    "password",
//     Addr:      "localhost:3306",
//     Net:       "tcp",
//     ParseTime: true,
//     Collation: "utf8mb4_unicode_ci",
//     Loc:       jst,
// }
// db, err := sql.Open("mysql", c.FormatDSN())

// 以下のようにDB接続を記載する方法もある。接続情報はこれがいい気がするが一旦は簡素化のためベタ書きする。
// user := os.Getenv("MYSQL_USER")
// pw := os.Getenv("MYSQL_PASSWORD")
// db_name := os.Getenv("MYSQL_DATABASE")
// var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
// var err error
// if Db, err = sql.Open("mysql", path); err != nil {
//     log.Fatal("Db open error:", err.Error())
// }
// checkConnect(100)

// fmt.Println("db connected!!")
