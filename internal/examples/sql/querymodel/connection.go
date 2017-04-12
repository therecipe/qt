package main

import (
	"github.com/therecipe/qt/sql"
	"github.com/therecipe/qt/widgets"
)

var db *sql.QSqlDatabase

func createConnection() bool {

	db = sql.QSqlDatabase_AddDatabase("QSQLITE", ":memory:")

	if !db.Open() {
		widgets.QMessageBox_Critical(nil, "Cannot open database", `Unable to establish a database connection.
This example needs SQLite support. Please read the Qt SQL driver documentation for information how to build it.

Click Cancel to exit.`, widgets.QMessageBox__Cancel, widgets.QMessageBox__NoButton)
		return false
	}

	var query = sql.NewQSqlQuery3(db)
	query.Exec("create table person (id int primary key, firstname varchar(20), lastname varchar(20))")
	query.Exec("insert into person values(101, 'Danny', 'Young')")
	query.Exec("insert into person values(102, 'Christine', 'Holand')")
	query.Exec("insert into person values(103, 'Lars', 'Gordon')")
	query.Exec("insert into person values(104, 'Roberto', 'Robitaille')")
	query.Exec("insert into person values(105, 'Maria', 'Papadopoulos')")

	query.Exec("create table items (id int primary key,imagefile int,itemtype varchar(20),description varchar(100))")
	query.Exec("insert into items values(0, 0, 'Qt','Qt is a full development framework with tools designed to streamline the creation of stunning applications and amazing user interfaces for desktop, embedded and mobile platforms.')")
	query.Exec("insert into items values(1, 1, 'Qt Quick','Qt Quick is a collection of techniques designed to help developers create intuitive, modern-looking, and fluid user interfaces using a CSS & JavaScript like language.')")
	query.Exec("insert into items values(2, 2, 'Qt Creator','Qt Creator is a powerful cross-platform integrated development environment (IDE), including UI design tools and on-device debugging.')")
	query.Exec("insert into items values(3, 3, 'Qt Project','The Qt Project governs the open source development of Qt, allowing anyone wanting to contribute to join the effort through a meritocratic structure of approvers and maintainers.')")

	query.Exec("create table images (itemid int, file varchar(20))")
	query.Exec("insert into images values(0, 'images/qt-logo.png')")
	query.Exec("insert into images values(1, 'images/qt-quick.png')")
	query.Exec("insert into images values(2, 'images/qt-creator.png')")
	query.Exec("insert into images values(3, 'images/qt-project.png')")

	return true
}
