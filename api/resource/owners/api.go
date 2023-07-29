package owners

import "database/sql"

type API struct {
	Database *sql.DB
}

func New(db *sql.DB) *API {
	return &API{
		Database: db,
	}
}
