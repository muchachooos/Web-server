package storage

import "main/model"

func (u *UserStorage) GetUserInDB(log, pass string) ([]model.Data, error) {

	var resultTable []model.Data

	err := u.DataBase.Select(&resultTable, "SELECT * FROM users WHERE login = ? AND password = ?", log, pass)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UserStorage) CreateUserInBD(log, pass string) error {

	_, err := u.DataBase.Exec("INSERT INTO users (login, password) VALUES (?,?)", log, pass)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStorage) DeleteUserFromDB(log, pass string) (bool, error) {

	res, err := u.DataBase.Exec("DELETE FROM users WHERE login = ? AND password = ?", log, pass)
	if err != nil {
		return false, err
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfDeletedRows == 0 {
		return false, err
	}

	return true, nil
}
