package orcldb

import (
	"database/sql"
	"fmt"

	//"os"

	"example/data"
)

var localDB = map[string]string{
	"service":  "XE",
	"username": "demo",
	"server":   "localhost",
	"port":     "1521",
	"password": "demo",
}

var autonomousDB = map[string]string{
	"service":        "MTHIS",
	"username":       "mthis",
	"server":         "192.168.222.150",
	"port":           "1521",
	"password":       "mthissatuduatiga",
	"walletLocation": ".",
}

func Dbmain() {
	db := GetSqlDBWithPureDriver(autonomousDB)
	//db := GetSqlDBWithGoDrOrDriver(autonomousDB)
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()
	sqlOperations(db)
	fmt.Println("DONE")

}

const createTableStatement = "CREATE TABLE TEMP_TABLE ( NAME VARCHAR2(100), CREATION_TIME TIMESTAMP DEFAULT SYSTIMESTAMP, VALUE  NUMBER(5))"
const dropTableStatement = "DROP TABLE TEMP_TABLE PURGE"
const insertStatement = "INSERT INTO TEMP_TABLE ( NAME , VALUE) VALUES (:name, :value)"
const queryStatement = "SELECT person_name, sex, birth_date from person where person_name = :name"

func sqlOperations(db *sql.DB) {
	selectPersonQuery, err := db.Prepare(queryStatement)
	handleError("prepare select person", err)
	defer selectPersonQuery.Close()

	person, err := selectPersonQuery.Query("RIVALDO")
	handleError("select person", err)
	defer person.Close()

	p := data.Person{}
	for person.Next(){
		person.Scan(&p.PersonName, &p.Sex, &p.BirthDate)
		fmt.Printf("%#v\n", p)
	}

	// _, err := db.Exec(createTableStatement)
	// handleError("create table", err)

	// defer db.Exec(dropTableStatement) // make sure the table is removed when all is said and done

	// stmt, err := db.Prepare(insertStatement)
	// handleError("prepare insert statement", err)

	// sqlresult, err := stmt.Exec("John", 42)
	// handleError("execute insert statement", err)

	// rowCount, _ := sqlresult.RowsAffected()
	// fmt.Println("Inserted number of rows = ", rowCount)

	// var queryResultName string
	// var queryResultTimestamp time.Time
	// var queryResultValue int32
	// row := db.QueryRow(queryStatement)
	// err = row.Scan(&queryResultName, &queryResultTimestamp, &queryResultValue)
	// handleError("query single row", err)
	// if err != nil {
	// 	panic(fmt.Errorf("error scanning db: %w", err))
	// }
	// fmt.Println(fmt.Sprintf("The name: %s, time: %s, value:%d ", queryResultName, queryResultTimestamp, queryResultValue))

	// _, err = stmt.Exec("Jane", 69)
	// handleError("execute insert statement", err)

	// _, err = stmt.Exec("Malcolm", 13)
	// handleError("execute insert statement", err)

	// // fetching multiple rows
	// theRows, err := db.Query(queryStatement)
	// handleError("Query for multiple rows", err)
	// defer theRows.Close()
	// var (
	// 	name  string
	// 	value int32
	// 	ts    time.Time
	// )
	// for theRows.Next() {
	// 	err := theRows.Scan(&name, &ts, &value)
	// 	handleError("next row in multiple rows", err)
	// 	fmt.Println(fmt.Sprintf("The name: %s and value:%d created at time: %s ", name, value, ts))
	// }
	// err = theRows.Err()
	// handleError("next row in multiple rows", err)
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		//os.Exit(1)
	}
}
