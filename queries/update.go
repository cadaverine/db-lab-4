package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func UpdateStudentPhone(db *sqlx.DB, id int, phone string) error {
	query := `
		UPDATE lms.students
		SET phone = $2
		WHERE id = $1
	`

	_, err := db.Exec(query, id, phone)
	if err != nil {
		fmt.Errorf("Unable to update student phone: %v", err)
	}

	return nil
}
