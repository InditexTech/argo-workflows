package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/upper/db/v4"
)

type DBType string

const (
<<<<<<< HEAD:persist/sqldb/db_type.go
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
	SQLite   dbType = "sqlite"
=======
	MySQL    DBType = "mysql"
	Postgres DBType = "postgres"
	SQLite   DBType = "sqlite"
>>>>>>> draft-3.6.5:util/sqldb/db_type.go
)

func DBTypeFor(session db.Session) DBType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}

func (t DBType) IntType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
