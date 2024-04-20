package main

import (
	"encoding/json"
	"fmt"
	dbtools "go-mysql-db/dbtool"
	"log"
	"os"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var conf Configuration

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	students := dbtools.SelectAllStudents()
	for _, student := range students {
		fmt.Println("ID:", student.ID, "\tName:", student.Name, "\tAge:", student.Age)
	}
}
