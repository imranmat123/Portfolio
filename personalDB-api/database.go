package personalDB_api

import (
	"ImransProfoiloWebsite/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func MynewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func CreateUser(d *sqlx.DB, user models.User) error {
	insert := `INSERT INTO  Users (
		first_name,
        last_name,
        email,
        password,
        profile_picture_url,
        about_me,
        created_at,
        updated_at,
        is_active
		) VALUES (?,?,?,?,?,?, NOW(),NOW(), TRUE)`

	_, err := d.Exec(insert,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.ProfilePicture,
		user.AboutMe,
	)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w ", err)
	}
	return nil
}

func UpdateUser(d *sqlx.DB, updatedUser models.User) error {

	if updatedUser.UserID == 0 {
		return fmt.Errorf("these two ID's do not match and thus are not to be updated")
	}
	update :=
		`UPDATE Users
		SET first_name = ?,
        last_name = ?,
        email = ?,
        password = ?,
        profile_picture_url = ?,
        about_me = ?,
        updated_at = NOW()
        WHERE user_id =?`

	_, err := d.Exec(
		update,
		updatedUser.FirstName,
		updatedUser.LastName,
		updatedUser.Email,
		updatedUser.Password,
		updatedUser.ProfilePicture,
		updatedUser.AboutMe,
		updatedUser.UserID,
	)

	if err != nil {
		return fmt.Errorf("Failed to update user: %w ", err)
	}
	return nil
}

func GetUser(d *sqlx.DB) ([]models.GetUser, error) {
	var u []models.GetUser
	g := `SELECT * FROM  Users`

	err := d.Select(&u, g)

	if err != nil {
		return nil, fmt.Errorf("sorry mate: %v", err)
	}
	return u, nil
}

func GetUserById(d *sqlx.DB, id int64) (models.GetUser, error) {
	var u = models.GetUser{}
	g := `SELECT * FROM Users WHERE user_id = ?`
	a := d.Get(&u, g, id)

	if a != nil {
		return u, fmt.Errorf("we haven't found the id: %v", a)
	} else {
		return u, nil
	}
}

func DeleteUser(d *sqlx.DB, id int64) error {
	var a = models.GetUser{}
	b := `DELETE FROM Users WHERE user_id = ?`

	c := d.Get(&a, b, id)

	if c != nil {
		return fmt.Errorf("item hasn't been deleted: %v", c)
	}
	return nil
}
