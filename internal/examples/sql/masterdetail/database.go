package main

import (
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"
)

var db *sql.QSqlDatabase

func createConnection() bool {

	db = sql.QSqlDatabase_AddDatabase("QSQLITE", "")
	db.SetDatabaseName(":memory:")
	if !db.Open() {
		widgets.QMessageBox_Critical(nil, "Cannot open database",
			`Unable to establish a database connection.
This example needs SQLite support. Please read
the Qt SQL driver documentation for information how
to build it.

Click Cancel to exit.`, widgets.QMessageBox__Cancel, 0)
		return false
	}

	query := sql.NewQSqlQuery3(db)

	query.Exec("create table artists (id int primary key, artist varchar(40), albumcount int)")

	query.Exec("insert into artists values(0, '<all>', 0)")
	query.Exec("insert into artists values(1, 'Ane Brun', 2)")
	query.Exec("insert into artists values(2, 'Thomas Dybdahl', 3)")
	query.Exec("insert into artists values(3, 'Kaizers Orchestra', 3)")

	query.Exec("create table albums (albumid int primary key, title varchar(50), artistid int, year int)")

	query.Exec("insert into albums values(1, 'Spending Time With Morgan', 1, 2003)")
	query.Exec("insert into albums values(2, 'A Temporary Dive', 1, 2005)")
	query.Exec("insert into albums values(3, '...The Great October Sound', 2, 2002)")
	query.Exec("insert into albums values(4, 'Stray Dogs', 2, 2003)")
	query.Exec("insert into albums values(5, 'One day you`ll dance for me, New York City', 2, 2004)")
	query.Exec("insert into albums values(6, 'Ompa Til Du D\xc3\xb8r', 3, 2001)")
	query.Exec("insert into albums values(7, 'Evig Pint', 3, 2002)")
	query.Exec("insert into albums values(8, 'Maestro', 3, 2005)")

	return true
}
