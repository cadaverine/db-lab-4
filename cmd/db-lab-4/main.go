package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/cadaverine/db-lab-4/queries"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lms_db"
)

func main() {
	dbURL := fmt.Sprintf("%s://%s:%v/%s", user, host, port, dbname)

	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	// Получаем студентов, закончивших курсы
	students, err := queries.SelectStudents(db, 2, 0)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	data, _ := json.Marshal(students)
	fmt.Printf("%+v\n\n", string(data))

	// Добавляем студента в базу
	student := &queries.Student{"Ragnar", "Lothbrok", "11-02-1990", "ragnar@lothbrok.com", ""}
	id, err := queries.InsertStudent(db, student)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("ID: %v\n\n", id)

	defer db.Close()
}
