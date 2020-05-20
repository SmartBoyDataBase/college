package model

import "sbdb-college/infrastructure"

type College struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Admin uint64 `json:"admin"`
}

func Get(id uint64) (College, error) {
	row := infrastructure.DB.QueryRow(`
	SELECT name, admin
	FROM college
	WHERE id=$1;
	`, id)
	result := College{
		Id: id,
	}
	err := row.Scan(&result.Name, &result.Admin)
	return result, err
}

func Create(college College) (College, error) {
	row := infrastructure.DB.QueryRow(`
	INSERT INTO college(name, admin)
	VALUES ($1, $2)
	RETURNING id;
	`, college.Name, college.Admin)
	err := row.Scan(&college.Id)
	return college, err
}

func Put(college College) (College, error) {
	_, err := infrastructure.DB.Exec(`
	UPDATE college
	SET name=$2,
	    admin=$3
	WHERE id=$1;
	`, college.Id, college.Name, college.Admin)
	return college, err
}

func Delete(id uint64) error {
	_, err := infrastructure.DB.Exec(`
	DELETE FROM college
	WHERE id=$1;
	`, id)
	return err
}

func All() ([]College, error) {
	rows, err := infrastructure.DB.Query(`
	SELECT (id, name, admin)
	FROM college;
	`)
	if err != nil {
		return nil, err
	}
	var result []College
	for rows.Next() {
		var college College
		err = rows.Scan(&college.Id, &college.Name, &college.Admin)
		if err != nil {
			return result, err
		}
		result = append(result, college)
	}
	return result, nil
}
