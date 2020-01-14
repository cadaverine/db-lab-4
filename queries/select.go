package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Student struct {
	Name   string
	Course string
}

func SelectStudents(db *sqlx.DB, limit, offset int) ([]Student, error) {
	query := `
		SELECT temp.name, c.name as course
		FROM lms.courses as c
		JOIN (
			SELECT CONCAT(first_name, ' ', last_name) as name, course_id
			FROM lms.students as s
			JOIN (
				SELECT student_id, course_id
				FROM lms.studying_processes
				WHERE status = 'done'
				ORDER BY receipt_date DESC
			) as p ON p.student_id = s.id
		) as temp ON temp.course_id = c.id
		LIMIT $1
		OFFSET $2;
	`

	var students []Student

	err := db.Select(&students, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch students: %v", err)
	}

	return students, nil
}
