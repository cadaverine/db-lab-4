package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DeleteStudent(db *sqlx.DB, id int) error {
	deleteProcessQuery := `
		DELETE FROM lms.studying_processes
		WHERE student_id = $1
	`

	deleteStudentQuery := `
		DELETE FROM lms.students
		WHERE id = $1
	`

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Unable to update student phone: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(deleteProcessQuery, id)
	if err != nil {
		return fmt.Errorf("Unable to delete studying process: %v", err)
	}

	_, err = tx.Exec(deleteStudentQuery, id)
	if err != nil {
		return fmt.Errorf("Unable to delete student: %v", err)
	}

	return nil
}
