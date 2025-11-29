package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type databaseType struct {
	*sql.DB
	Path             string
	ConnectionString string
}

func (database *databaseType) InitializeDatabase() error {
	dbDir := "DBs"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("failed to create DBs directory: %v", err)
	}

	dbPath := filepath.Join(dbDir, "DB.db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("Database not found. Creating new database...")

		file, err := os.Create(dbPath)
		if err != nil {
			return fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close()

		fmt.Println("Database file created successfully")
	} else {
		fmt.Println("Database already exists")
	}

	database.Path = dbPath
	return nil
}

func (database *databaseType) executeDDL() error {
	ddlStatements := []string{
		"DROP TABLE IF EXISTS PROJECT;",
		"CREATE TABLE IF NOT EXISTS PROJECT (id integer primary key autoincrement, name varchar not null);",
		"DROP TABLE IF EXISTS ITEM;",
		"CREATE TABLE IF NOT EXISTS ITEM (id integer primary key autoincrement, id_project integer not null, name varchar not null, comment varchar, FOREIGN KEY (id_project) REFERENCES project (id) ON DELETE CASCADE);",
		"DROP TABLE IF EXISTS SLIDE;",
		"CREATE TABLE IF NOT EXISTS SLIDE (id integer primary key autoincrement, id_item integer not null, num integer, name text, content text, content_type text not null default 'code', direct text not null default 'column', content_proportion integer not null default 1, page_proportion integer not null default 2, comment text, FOREIGN KEY (id_item) REFERENCES item (id) ON DELETE CASCADE);",
		"DROP TABLE IF EXISTS TAB;",
		"CREATE TABLE IF NOT EXISTS TAB (id integer primary key autoincrement, id_slide integer not null, num integer, name text, content text, content_type text not null default 'table', comment text, FOREIGN KEY (id_slide) REFERENCES slide (id) ON DELETE CASCADE);",
	}

	for _, statement := range ddlStatements {
		if strings.TrimSpace(statement) == "" {
			continue
		}

		_, err := database.Exec(statement)
		if err != nil {
			return fmt.Errorf("failed to execute statement '%s': %v", statement, err)
		}
	}

	return nil
}

var database databaseType

func (database *databaseType) buildConnectionString() {
	database.ConnectionString = database.Path
}

func (database *databaseType) Connect() error {
	if err := database.InitializeDatabase(); err != nil {
		return fmt.Errorf("database initialization failed: %v", err)
	}

	db, err := sql.Open("sqlite", database.Path)
	check(err)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	database.DB = db

	if database.needsDDL() {
		fmt.Println("Running DDL to create tables...")
		if err := database.executeDDL(); err != nil {
			return fmt.Errorf("failed to execute DDL: %v", err)
		}
		fmt.Println("Database initialized successfully")
	} else {
		fmt.Println("Database tables already exist")
	}

	return nil
}

func (database *databaseType) needsDDL() bool {
	query := "SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'"

	var count int
	err := database.QueryRow(query).Scan(&count)
	if err != nil {
		return true
	}

	return count == 0
}

func check(err interface{}) {
	if err != nil {
		_, fileName, lineNo, _ := runtime.Caller(1) // Получаем информацию о вызывающем файле
		log.Printf("%s: %d\n", filepath.Base(fileName), lineNo)
		log.Println(err)

		panic(err)
	}
}

func getContent(idItem int) []slide {
	resCont := make([]slide, 0)

	query := `select id
     			   , id_item
    			   , num
     			   , name
    			   , content
    			   , content_type
    			   , direct
				   , content_proportion
				   , page_proportion
				   , comment
				from slide
			  where id_item = ` + strconv.Itoa(idItem) + ` order by num asc`

	cont, err := database.Query(query)
	check(err)

	var id, id_item, num, contentProportion, pageProportion sql.NullInt64
	var name, content_, contentType, direct, comment sql.NullString
	for cont.Next() {
		err = cont.Scan(&id, &id_item, &num, &name, &content_, &contentType, &direct, &contentProportion, &pageProportion, &comment)
		check(err)

		tabs := getTabs(int(id.Int64))
		resCont = append(resCont, slide{int(id.Int64), idItem, int(num.Int64), name.String, content_.String, contentType.String, direct.String, int(contentProportion.Int64), int(pageProportion.Int64), comment.String, tabs, 0})
	}

	err = cont.Close()
	check(err)

	return resCont
}

func getItemId(itemName string, bGetFirst bool) int {

	query := `select id
				   , name
				from item
			  where lower(name) like lower('%` + itemName + `%')`

	cont, err := database.Query(query)
	check(err)

	var id sql.NullInt64
	var name sql.NullString
	for cont.Next() {
		err = cont.Scan(&id, &name)
		check(err)

		if bGetFirst {
			return int(id.Int64)
		} else {
			fmt.Println(name.String)
		}
	}

	return -1
}

func countItems(itemName string) int {

	query := `select count(1) cnt
				from item
			  where lower(name) like lower('%` + itemName + `%')`

	cont, err := database.Query(query)
	check(err)

	var count sql.NullInt64
	cont.Next()
	err = cont.Scan(&count)
	check(err)

	return int(count.Int64)
}

func getTabs(idSlide int) []Tab {

	resTabs := make([]Tab, 0)

	query := `select id
     			   , id_slide
    			   , num
     			   , name
    			   , content
    			   , content_type
				   , comment
				from tab
			  where id_slide = ` + strconv.Itoa(idSlide) + ` order by num asc`

	tabs, err := database.Query(query)
	check(err)

	var id, id_slide, num sql.NullInt64
	var name, content_, contentType, comment sql.NullString
	for tabs.Next() {
		err = tabs.Scan(&id, &id_slide, &num, &name, &content_, &contentType, &comment)
		check(err)

		resTabs = append(resTabs, Tab{int(id.Int64), idSlide, int(num.Int64), name.String, content_.String, contentType.String, comment.String})
	}

	err = tabs.Close()
	check(err)

	return resTabs
}
