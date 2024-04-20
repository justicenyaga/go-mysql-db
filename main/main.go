package main

import (
	"encoding/json"
	"fmt"
	dbtools "go-mysql-db/dbtool"
	"go-mysql-db/models"
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

	student := models.Student{
		ID:   3,
		Name: "Robert",
		Age:  28,
	}
	lastInsertedID := dbtools.Save(student)
	fmt.Println("Last Inserted ID", lastInsertedID)
}
