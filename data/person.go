package data

import (
	"encoding/json"
	"io"
)

type Person struct {
	PersonId	int64	`db:"person_id"`
	PersonName 	string	`db:"person_name"`
	Sex      	string	`db:"sex"`
	BirthDate   string	`db:"birth_date"`
	IdNo		string 	`db:"id_no"`
}

type Persons []*Person

func (p *Persons) ToJSON(w io.Writer) error{
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Persons) FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}



