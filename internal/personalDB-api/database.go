package personalDB_api

import (
	"ImransProfoiloWebsite/internal/models"
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

//users

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

func UpdateUser(d *sqlx.DB, updatedUser models.GetUser) error {

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

// projects
func CreateProjects(db *sqlx.DB, project models.Projects) error {
	insert := `INSERT INTO Project(
        user_id,            
        project_name,
        about,  
        created_at,  
        updated_at,
        project_url) VALUES (?, ?, ?, NOW(), NOW(), ?)`

	_, err := db.Exec(insert,
		project.UserID,
		project.ProjectName,
		project.About,
		project.ProjectURL,
	)
	if err != nil {
		return fmt.Errorf("CreateProjects has failed in insert: %v", err)
	}
	return nil
}

func GetAllProjects(db *sqlx.DB) ([]models.Projects, error) {

	var p []models.Projects

	get := `SELECT * FROM Project`

	err := db.Select(&p, get)
	if err != nil {
		return nil, fmt.Errorf("we are unable to get all projects: %v", err)
	}
	return p, nil
}

func GetProjectById(db *sqlx.DB, id int64) (*models.Projects, error) {
	var a models.Projects
	in := `SELECT * FROM Project WHERE project_id = ?`

	err := db.Get(&a, in, id)
	if err != nil {
		return nil, fmt.Errorf("we have not been able to get the project by the id: %v", err)
	}
	return &a, nil
}

func UpdateProjects(db *sqlx.DB, m *models.Projects) (*models.Projects, error) {
	up := `UPDATE Project SET 
                project_name =?,
                about =?,
                updated_at =?,
            	project_url=?
                WHERE project_id = ?`

	_, err := db.Exec(up,
		m.ProjectName,
		m.About,
		m.UpdatedAt,
		m.ProjectURL,
		m.ProjectID,
	)

	if err != nil {
		return nil, fmt.Errorf("we have not been able to update the project: %v", err)
	}

	a, err1 := GetProjectById(db, m.ProjectID)

	if err1 != nil {
		return nil, fmt.Errorf("we have atampted to get the ID from the update function, it was unable to be found: %v", err1)
	}

	return a, nil
}

func DeleteProject(db *sqlx.DB, id int64) error {
	del := `DELETE FROM Project WHERE project_id =?`

	_, err := db.Exec(del, id)

	if err != nil {
		return fmt.Errorf("we have not been able to delete the project: %v", err)
	}
	return nil
}
