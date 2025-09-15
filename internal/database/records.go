package database

import "time"

type Record struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []Tag     `json:"tags"`
}

func GetRecords(user *User) ([]Record, error) {
	tags := []Tag{}
	rows, err := db.Query("SELECT id, name FROM tags WHERE user_login = $1")

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

func CreateRecord(userLogin string, content string, tagIds []int64) (recordId int64, err error) {
	row := db.QueryRow(
		"INSERT INTO records(user_login, content) VALUES ($1, $2) RETURNING id",
		userLogin,
		content,
	)
	err = row.Scan(&recordId)
	if err != nil {
		return
	}

	err = UpdateRecordTags(recordId, tagIds)
	if err != nil {
		return
	}
	
	return
}

func UpdateRecordContent(recordId int64, content string) error {
	// TODO
	return nil
}

func UpdateRecordTags(recordId int64, tagIds []int64) error {
	return nil
}

func DeleteRecord(id int64) error {
	_, err := db.Exec("DELETE FROM records WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
