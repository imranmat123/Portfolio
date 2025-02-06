package Types

import "github.com/jmoiron/sqlx"

type ConnectToDatabase struct {
	D *sqlx.DB
}
