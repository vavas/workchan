package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/vavas/workchan/app/config"
	"log"

	_ "github.com/lib/pq"
)

func Load(conf *config.Database) (*DB, error) {
	masterDB, err := connect(conf.MasterURL, conf.PoolSize)
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to the master DB")
	}
	readerDB, err := connect(conf.ReaderURL, conf.PoolSize)
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to the master DB")
	}
	db := &DB{
		Master: masterDB,
		Reader: readerDB,
	}

	return db, nil

}

func connect(dbURL string, poolSize int) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(poolSize)
	return db, nil
}
