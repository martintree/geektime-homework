package dao

import "github.com/pkg/errors"

type User struct {
	ID int
	name string
}

func QueryUser(ID int) (*User, error) {
	row := MysqlDB.QueryRow("select id from o_user where id = ?", ID)
	qUser := User{}
	if err := row.Scan(&qUser.ID); err != nil {
		return nil, errors.Wrap(err, "query user occur error")
	}
	return &qUser, nil
}