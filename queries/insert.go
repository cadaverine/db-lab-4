package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Student struct {
	FirstName string
	LastName  string
	BirthDate string
	Email     string
	Phone     string
}

func InsertStudent(db *sqlx.DB, s *Student) (int64, error) {
	query := `
		INSERT INTO lms.students (
			first_name,
			last_name,
			phone,
			email,
			birthdate
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var id int64
	err := db.QueryRow(query, s.FirstName, s.LastName, s.Phone, s.Email, s.BirthDate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Unable to insert student: %v", err)
	}

	return id, nil
}
