package database

import "time"

type Record struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetRecords(user *User) ([]Record, error) {
	rows, err := db.Query(
		"SELECT id, content, created_at FROM records WHERE user_login = $1",
		user.Login,
	)
	if err != nil {
		return nil, err
	}

	result := []Record{}
	for rows.Next() {
		record := Record{}
		err = rows.Scan(&record.Id, &record.Content, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func CreateRecord(user *User, content string) (*Record, error) {
	row := db.QueryRow(
		"INSERT INTO records(user_login, content) VALUES ($1, $2) RETURNING id, content, created_at",
		user.Login,
		content,
	)
	record := &Record{}
	err := row.Scan(&record.Id, &record.Content, &record.CreatedAt)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func DeleteRecord(id int64) error {
	_, err := db.Exec("DELETE FROM records WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
