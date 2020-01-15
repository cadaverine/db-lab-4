package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DeleteStudent(db *sqlx.DB, id int) error {
	query := `
		DELETE FROM lms.students
		WHERE id = $1
	`

	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Errorf("Unable to delete student: %v", err)
	}

	return nil
}
