package main

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type databaseType struct {
	*sql.DB
	Path             string
	ConnectionString string
}

var database databaseType

func (database *databaseType) buildConnectionString() {
	database.ConnectionString = database.Path
}

func (database *databaseType) Connect() error {
	db, err := sql.Open("sqlite", "DBs"+string(os.PathSeparator)+"DB.db")
	check(err)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	database.DB = db
	return nil
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
