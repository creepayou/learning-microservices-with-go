package orcldb

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

var dbParams = map[string]string{
	"dbname":        "tubandev",
	"username":       "postgres",
	"password":       "G00gl3123",
	"host":         "192.168.222.90",
	"port":           "5432",
	"walletLocation": ".",
}

func GetSqlDBWithPureDriver() *sql.DB {
	connectionString := "postgres://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["host"] + ":" + dbParams["port"] + "/" + dbParams["dbname"] + "?sslmode=disable"
	// if val, ok := dbParams["walletLocation"]; ok && val != "" {
	// 	connectionString += "?SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(dbParams["walletLocation"])
	// }
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}
	return db
}
