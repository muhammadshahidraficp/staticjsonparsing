package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "vinam"
	dbPass := "vinam"
	dbName := "goblog"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var JSON = `{
    "name":"Mark Taylor",
    "jobtitle":"Software Developer",
    "phone":{
        "home":"123-466-799",
        "office":"564-987-654"
    },
    "email":"markt@gmail.com"
}`

func main() {
	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)
	name := info["name"]
	jobtitle := info["jobtitle"]
	email := info["email"]
	home := info["phone"].(map[string]interface{})["home"]
	office := info["phone"].(map[string]interface{})["office"]
	fmt.Println(name)
	fmt.Println(jobtitle)
	fmt.Println(email)
	fmt.Println(home)
	fmt.Println(office)
	db := dbConn()
	insForm, err := db.Prepare("INSERT INTO employee(name,jobtitle,email,home,office) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(name, jobtitle, email, home, office)
}
