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

func Save(student models.Student) int64 {
	db := connect()
	defer db.Close()

	save, err := db.Prepare("insert into student (id, name, age) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := save.Exec(student.ID, student.Name, student.Age)
	if err != nil {
		log.Fatal(err.Error())
	}

	studentID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	return studentID
}

func Update(student models.Student) int64 {
	db := connect()
	defer db.Close()

	update, err := db.Prepare("update student set name = ?, age = ? where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := update.Exec(student.Name, student.Age, student.ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}

	return rowsAffected
}
