package migration

import (
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
)

type Repository struct {
	Conn *mysql.Conn
}

func NewRepository(conn *mysql.Conn) *Repository {
	return &Repository{Conn: conn}
}

func (r *Repository) Execute(sql string, args ...interface{}) (middleware.Result, error) {
	return r.Conn.Execute(sql, args...)
}
