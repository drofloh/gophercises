package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "andrewholford"
	password = "1234qwer"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	err = resetDB(db, dbname)
	must(err)

	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	must(err)

	defer db.Close()

	must(createPhoneNumbersTable(db))
	_, err = insertPhone(db, "1234567890")
	must(err)
	_, err = insertPhone(db, "123 456 7891")
	must(err)
	_, err = insertPhone(db, "(123) 456 7892")
	must(err)
	_, err = insertPhone(db, "(123) 456-7893")
	must(err)
	_, err = insertPhone(db, "123-456-7894")
	must(err)
	_, err = insertPhone(db, "123-456-7890")
	must(err)
	_, err = insertPhone(db, "1234567892")
	must(err)
	_, err = insertPhone(db, "(123)456-7892")
	must(err)

	phones, err := allPhones(db)
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on ... %+v\n", p)
		number := normalize(p.number)
		if number != p.number {
			fmt.Println("Updating or removing ", number)
			existing, err := findPhone(db, number)
			must(err)
			if existing != nil {
				//delete
				must(deletePhone(db, p.id))
			} else {
				//update
				p.number = number
				must(updatePhone(db, p))
			}
		} else {
			fmt.Println("No changes required")

		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type phone struct {
	id     int
	number string
}

func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id).Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", number).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func allPhones(db *sql.DB) ([]phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []phone

	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING ID`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
          CREATE TABLE IF NOT EXISTS phone_numbers (
            id SERIAL,
            value VARCHAR(255)
          )
        `
	_, err := db.Exec(statement)
	return err
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, dbname)
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// regex version of above
//func normalize(phone string) string {
//  re := regexp.MustCompile("\\D")
//  return re.ReplaceAllString(phone, "")
//}
