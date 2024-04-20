package dbtools

import (
	"database/sql"
	"go-mysql-db/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driverName     string
	dataSourceName string
)

func DBInitializer(dn, dsn string) {
	driverName = dn
	dataSourceName = dsn
}

func connect() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func SelectAllStudents() []models.Student {
	db := connect()
	defer db.Close()

	rows, err := db.Query("select * from student")
	if err != nil {
		log.Fatal(err.Error())
	}

	students := []models.Student{}

	for rows.Next() {
		student := models.Student{}

		err = rows.Scan(&student.ID, &student.Name, &student.Age)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		students = append(students, student)
	}

	return students
}

func SelectStudentByID(id int) models.Student {
	db := connect()
	defer db.Close()

	row := db.QueryRow("select * from student where id = ?", id)

	student := models.Student{}

	err := row.Scan(&student.ID, &student.Name, &student.Age)
	if err != nil {
		log.Fatal(err.Error())
	}

	return student
}
