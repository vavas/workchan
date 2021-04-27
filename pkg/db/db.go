package db

import "github.com/jmoiron/sqlx"

type DB struct {
	Master *sqlx.DB
	Reader *sqlx.DB
}

func (d *DB) Close() {
	if d.Master.DB == nil {
		d.Master.Close()
	}

	if d.Reader.DB == nil {
		d.Master.Close()
	}
}
