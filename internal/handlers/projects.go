package handlers

import (
	"ImransProfoiloWebsite/internal/models"
	personalDB_api "ImransProfoiloWebsite/internal/personalDB-api"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ProjectHandler(db *sqlx.DB, c echo.Context) error {
	if db != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to connect to the database",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "we are connected to the database",
	})
}

func CreateProject(db *sqlx.DB, c echo.Context) error {
	var a models.Projects
	err := c.Bind(&a)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "we are unable to bind the new project",
		})
	}
	u := personalDB_api.CreateProjects(db, a)
	if u != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to create project",
		})
	}
	c.JSON(201, map[string]interface{}{
		"status": 201,
		"data":   u,
	})
	return nil
}

func GetAllProjects(db *sqlx.DB, c echo.Context) error {

	a, err := personalDB_api.GetAllProjects(db)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to get all projects form the database",
		})
	}

	if len(a) == 0 {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "query was successful, but there is no projects",
			"data":    a,
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   a,
	})
}

func GetProjectById(db *sqlx.DB, c echo.Context) error {
	id, iderr := strconv.Atoi(c.Param("id"))
	if iderr != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to convert the ID",
		})
	}

	b, err1 := personalDB_api.GetProjectById(db, int64(id))
	if err1 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  404,
			"message": "There is no matching ProjectID",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})

}

func GetByABoutMe(db *sqlx.DB, c echo.Context) error {

	a, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to convert ID",
		})
	}

	b, err1 := personalDB_api.GetProjectById(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":   404,
			"messsage": "unable to find project",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b.About,
	})
}

func GetProjectURL(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to covert the ID",
		})
	}
	d, err1 := personalDB_api.GetProjectById(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find project",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   d.ProjectURL,
	})
}
func GetProjectName(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "unable to bind to variable",
		})
	}
	d, err1 := personalDB_api.GetProjectById(db, int64(a))
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to find project",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   d.ProjectName,
	})
}

func UpdateProject(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  404,
			"message": "unable to convert ID",
		})
	}
	get, Gerr := personalDB_api.GetProjectById(db, int64(a))

	if Gerr != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  404,
			"message": "ID does not match the ID in the database",
		})
	}

	b, err1 := personalDB_api.UpdateProjects(db, get)
	if err1 != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  50,
			"message": "unable to update project",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   b,
	})
}

func DeleteProject(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"status":  404,
			"message": "unable to bind variable",
		})
	}
	b := personalDB_api.DeleteProject(db, int64(a))
	if b != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "was unable to delete the project",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "project has been deleted",
	})
}

func GetProjectsByUserID(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to convert the ID",
		})
	}
	all, err1 := personalDB_api.GetAllProjects(db)
	if err1 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to convert the ID",
		})
	}

	var theProject []models.Projects

	for i := 0; i < len(all); i++ {
		if all[i].UserID == int64(a) {
			theProject = append(theProject, all[i])
		}
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   len(theProject),
	})
}

func GetAllUserProjects(db *sqlx.DB, c echo.Context) error {
	a, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to convert the ID",
		})
	}
	b, err1 := personalDB_api.GetAllProjects(db)
	if err1 != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "we are unable to convert the ID",
		})
	}

	var projects []models.Projects

	for i := 0; i < len(b); i++ {
		if b[i].UserID == int64(a) {
			projects = append(projects, b[i])
		}
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   projects,
	})

}
