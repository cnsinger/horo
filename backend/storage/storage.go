package storage

import (
	"database/sql"
	"fmt"
	"horo/model"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const timerSql = `CREATE TABLE IF NOT EXISTS timer (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"context" TEXT,
		"insertAt" INTEGER,
		"length" INTEGER,
		"doneAt" INTEGER,
		"done" INTEGER
	  );`
const dbPath = "./sqlite-database.db"

type Database interface {
	CreateTable(string)
	//Insert(string)
	InsertTimer(model.HoroTimer)
	Query() []*model.HoroTimer
	Delete()
	Update([]int)
}

var defaultDatabase Database

type Storage struct {
	db *sql.DB
}

// TODO: 修改为单例模式
func Instance() Database {
	if defaultDatabase != nil {
		return defaultDatabase
	}
	defaultDatabase = InitDatabase(dbPath)
	return defaultDatabase
}

func InitDatabase(path string) Database {
	_, exists := os.Stat(path)
	if os.IsNotExist(exists) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}
	sqliteDatabase, _ := sql.Open("sqlite3", path)
	//defer sqliteDatabase.Close()

	s := Open(sqliteDatabase)
	s.CreateTable(timerSql)
	return s
}

func Open(db *sql.DB) Database {
	return &Storage{db}
}

func (s *Storage) CreateTable(timerSql string) {
	statement, err := s.db.Prepare(timerSql) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func (s *Storage) Insert(values string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := s.db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(values)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (s *Storage) InsertTimer(timer model.HoroTimer) {
	tSql := "INSERT INTO timer(context, insertAt, length, doneAt, done) VALUES (?, ?, ?, ?, 0)"
	statement, err := s.db.Prepare(tSql)
	if err != nil {
		log.Fatalln("prepare: ", err.Error())
	}
	_, err = statement.Exec(timer.Context, timer.InsertAt.Unix(), timer.Length, timer.InsertAt.Unix()+int64(timer.Length))
	if err != nil {
		log.Fatalln("insert: ", err.Error())
	}
}

func (s *Storage) Query() []*model.HoroTimer {
	cSql := "SELECT id, context, insertAt, length, doneAt FROM timer WHERE done != 1;"
	row, err := s.db.Query(cSql)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	horoTimers := []*model.HoroTimer{}
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var context string
		var insertAt int
		var length int
		var doneAt int
		row.Scan(&id, &context, &insertAt, &length, &doneAt)
		t := model.HoroTimer{Id: id, Context: context, InsertAt: time.Unix(int64(insertAt), 0), Length: length, DoneAt: time.Unix(int64(doneAt), 0)}
		horoTimers = append(horoTimers, &t)
	}
	return horoTimers
}

func (s *Storage) Delete() {
	dSql := "DELETE FROM timer where done = 1"
	_, err := s.db.Exec(dSql)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (s *Storage) Update(ids []int) {
	sIds := []string{}
	for i := range ids {
		sIds = append(sIds, strconv.Itoa(i))
	}
	uSql := fmt.Sprintf("UPDATE timer set done = 1 where id in (%s) ", strings.Join(sIds, ","))
	_, err := s.db.Exec(uSql)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
