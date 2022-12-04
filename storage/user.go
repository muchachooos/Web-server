package storage

import "Web-server/model"

func (u *UserStorage) GetLoginUserInDB(log, pass string) ([]model.Data, error) {

	var resultTable []model.Data

	err := u.DataBase.Select(&resultTable, "SELECT * FROM users WHERE login = ? AND password = ?", log, pass)
	if err != nil {
		return nil, err
	}

	return resultTable, nil
}

func (u *UserStorage) RegistrationUserInBD(log, pass string) error {

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

func (u *UserStorage) ChangePassUserInDB(log, pass, newPass string) (bool, error) {

	var resultTable []model.Data

	err := u.DataBase.Select(&resultTable, "SELECT * FROM users WHERE login = ? AND password = ?", log, pass)
	if err != nil {
		return false, err
	}

	if len(resultTable) == 0 {
		return false, err
	}

	res, err := u.DataBase.Exec("UPDATE users SET password = ? WHERE login = ? AND password = ?", newPass, log, pass)
	if err != nil {
		return false, err
	}

	countOfChangedRows, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if countOfChangedRows == 0 {
		return false, err
	}

	return true, nil
}
