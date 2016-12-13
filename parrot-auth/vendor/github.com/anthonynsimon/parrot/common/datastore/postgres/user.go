package postgres

import "github.com/anthonynsimon/parrot/common/model"

func (db *PostgresDB) GetUserByEmail(email string) (*model.User, error) {
	u := model.User{}
	row := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, parseError(err)
	}

	return &u, nil
}

func (db *PostgresDB) GetUserByID(id string) (*model.User, error) {
	u := model.User{}
	row := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, parseError(err)
	}

	return &u, nil
}

func (db *PostgresDB) CreateUser(u model.User) (*model.User, error) {
	row := db.QueryRow("INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id, name, email", u.Name, u.Email, u.Password)
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	return &u, parseError(err)
}
