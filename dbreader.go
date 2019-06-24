package dbreader

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type DBReader interface {
	Read() (interface{}, error)
	Write() (interface{}, error)
	Close() (error)
}

type Reader struct {
	DB *sql.DB
}

func NewReader(username, password, database string) (*Reader, error) {
	db, err := sql.Open("mysql", username + ":" + password + "@/" + database)
	if err != nil {
		return nil, err
	}
	return &Reader{
		db,
	}, nil
}

func (this *Reader) Read(statement string) (*sql.Rows, error) {
	rows, err := this.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (this *Reader) Write(statement string) (bool, error) {
	_, err := this.DB.Exec(statement)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *Reader) Close() error {
	return this.DB.Close()
}