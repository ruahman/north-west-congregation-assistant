package models

import "database/sql"

func Map[T interface{}](rows *sql.Rows, f func(*sql.Rows) (T, error)) ([]T, error) {
	defer rows.Close()
	var result []T
	for rows.Next() {
		r, err := f(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
