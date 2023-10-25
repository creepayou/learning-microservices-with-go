package handlers

import (
	"database/sql"
	"example/data"
	"fmt"
	"log"
	"net/http"

	pgdb "example/pg-db"

	_ "github.com/lib/pq"
)

type PersonHandler struct {
	l *log.Logger
}

func NewPersonHandler(l *log.Logger) *PersonHandler{
	return &PersonHandler{l}
}

func (p*PersonHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodGet{
		// p.GetPerson(rw, r)
		p.GetPerson(rw, r)
		return
	}		

	// if r.Method == http.MethodPut{
	// 	p.updateProduct(rw, r)
	// 	return
	// }	

	//catchAll
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*PersonHandler) GetPerson(rw http.ResponseWriter, r *http.Request){
	db := getDbConnection()
	defer closeDbConnection(db)

	var name string = r.URL.Query().Get("name")
	fmt.Println("Name to search:", name)
	selectPersonQuery := "SELECT person_id, person_name, sex, birth_date, id_no from person where person_name LIKE '%' || $1 || '%'"

	selectPersonPrep, err := db.Prepare(selectPersonQuery)
	handleError("Preparing Select Person Query.", err)

	rows, err := selectPersonPrep.Query(name)
	handleError("Running Select Person Query.", err)
	defer rows.Close()

	person := data.Person{}
	for rows.Next(){
		rows.Scan(&person.PersonId, &person.PersonName, &person.Sex, &person.BirthDate, &person.IdNo)
		fmt.Printf("%#v\n", person)
	}
}

func getDbConnection() *sql.DB{
	db := pgdb.GetSqlDBWithPureDriver()
	//db := GetSqlDBWithGoDrOrDriver(autonomousDB)
	fmt.Println("Connection Established.\n====================")
	return db
}

func closeDbConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		//os.Exit(1)
	}
}


