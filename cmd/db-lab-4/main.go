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
		fmt.Fprintf(os.Stderr, "Database connection failed: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Database is anavailable: %v\n", err)
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

	// Обновляем номер телефона студента
	err = queries.UpdateStudentPhone(db, 10, "89999999999")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Удаляем студента и его учебный прогресс (транзакция)
	err = queries.DeleteStudent(db, 9)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
