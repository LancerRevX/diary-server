package database

import "time"

type Record struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      []Tag     `json:"tags"`
}

func GetRecordById(recordId int64) (*Record, error) {
	row := db.QueryRow(
		"SELECT id, content, created_at FROM records WHERE id = $1",
		recordId,
	)
	record := Record{}
	err := row.Scan(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func GetRecords(userId int64) ([]Record, error) {
	rows, err := db.Query(
		"SELECT id, content, created_at FROM records WHERE user_id = $1",
		userId,
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

func CreateRecord(userId int64, content string, tagIds []int64) (recordId int64, err error) {
	row := db.QueryRow(
		"INSERT INTO records(user_id, content) VALUES ($1, $2) RETURNING id",
		userId,
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

func RecordExists(recordId int64, userId int64) (bool, error) {
	row := db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM records WHERE id = $1 AND user_id = $2)",
		recordId,
		userId,
	)
	var result bool
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}

func DeleteRecord(id int64) error {
	_, err := db.Exec("DELETE FROM records WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
