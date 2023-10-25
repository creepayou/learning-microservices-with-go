package orcldb

import (
	"fmt"

	//"os"
	"example/data"

	_ "github.com/lib/pq"
)

var localDB = map[string]string{
	"service":  "XE",
	"username": "demo",
	"server":   "localhost",
	"port":     "1521",
	"password": "demo",
}

var autonomousDB = map[string]string{
	"dbname":        "tubandev",
	"username":       "postgres",
	"password":       "G00gl3123",
	"host":         "192.168.222.90",
	"port":           "5432",
	"walletLocation": ".",
}

const queryStatement = "SELECT person_name, sex, birth_date from person where person_name = $1"

func Dbmain() {

	conn := GetSqlDBWithPureDriver()
	//db := GetSqlDBWithGoDrOrDriver(autonomousDB)
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	selectPersonQuery, err := conn.Prepare(queryStatement)
	handleError("Preparing Select Person Query.", err)

	rows, err := selectPersonQuery.Query("SUISUI")
	handleError("Running Select Person Query.", err)
	defer rows.Close()

	p := data.Person{}
	for rows.Next(){
		rows.Scan(&p.PersonName, &p.Sex, &p.BirthDate)
		fmt.Printf("%#v\n", p)
	}

}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		//os.Exit(1)
	}
}
