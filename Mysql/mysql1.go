package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "Test"
	db       *sql.DB
)

type studentInfo struct {
	Name  string
	Age   int32
	Class string
	Grade int32
	Sex   string
}

func main() {
	info := studentInfo{
		Name:  "name",
		Age:   22,
		Class: "",
		Grade: 4,
		Sex:   "男",
	}
	age, err := getAge(info.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			str := "INSERT INTO `student`(`name`, `age`, `class`, `grade`, `sex`) values('%s', %d, '%s', %d, '%s')"
			newSql := fmt.Sprintf(str, info.Name, info.Age, info.Class, info.Grade, info.Sex)
			fmt.Println("insert into sql: ", newSql)
			affect, err := exec(newSql)
			if err != nil {
				fmt.Println("exec error: ", err)
				return
			}
			fmt.Println("exec success, affect: ", affect)
			age, err = getAge(info.Name)
			if err != nil {
				fmt.Println("getAge error: ", err)
				return
			}
			fmt.Println("getAge success, age: ", age)
			return
		} else {
			fmt.Println("getAge error: ", err)
			return
		}
	}
	fmt.Println("getAge success, age: ", age)
	//db, err := sql.Open("mysql", strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, ""))
	//if err != nil {
	//	fmt.Println("open error: ", err)
	//} else {
	//	fmt.Println("open success")
	//}
	//defer db.Close()
	//fmt.Println("关闭db")
}

func initMysql() (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, ""))
		if err != nil {
			return db, fmt.Errorf("connect to mysql failed: %s", err)
		}
		err = db.Ping()
		if err != nil {
			return db, fmt.Errorf("ping mysql failed: %s", err)
		}
		db.SetMaxOpenConns(100) // 最大打开连接数
		db.SetMaxIdleConns(100) // 最大空闲连接数
	}
	return db, nil
}

func getAge(name string) (age int32, err error) {
	str := "SELECT `age` FROM student WHERE `name` = '%s'"
	newSql := fmt.Sprintf(str, name)
	fmt.Println("select sql: ", newSql)
	db, err := initMysql()
	if err != nil {
		return 0, err
	}
	row := db.QueryRow(newSql)
	err = row.Scan(&age)
	return
}

func exec(str string) (int64, error) {
	db, err := initMysql()
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare(str)
	if err != nil {
		return 0, fmt.Errorf("db.Prepare error: %s", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec()
	if err != nil {
		return 0, fmt.Errorf("stmt.Exec error: %s", err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("res.RowAffected error: %s", err)
	}
	return affect, nil
}
