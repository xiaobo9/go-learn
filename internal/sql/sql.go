// https://www.sqlitetutorial.net/sqlite-go/connect/
package sql

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"

	"log"

	_ "github.com/glebarez/go-sqlite"
)

type Country struct {
	Id         int
	Name       string
	Population int
	Area       int
}

func Demo() {
	// db, err := sql.Open("sqlite", ":memory:")
	// db, err := sql.Open("sqlite", "./my.db")
	// "./my.db?_pragma=foreign_keys(1)"
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to the SQLite database successfully.")

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(sqliteVersion)

	// create the countries table
	_, err = CreateTable(db)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Table countries was created successfully.")

	// create a new country
	country := &Country{
		Name:       "United States",
		Population: 329064917,
		Area:       9826675,
	}
	// insert the country
	countryId, err := Insert(db, country)
	if err != nil {
		log.Println(err)
		return
	}

	// print the inserted country
	log.Printf(
		"The country %s was inserted with ID:%d\n",
		country.Name,
		countryId,
	)

	// Update the population of a country
	_, err = Update(db, 1, 346037975)
	if err != nil {
		log.Println(err)
		return
	}

	// Delete the country with id 1
	_, err = Delete(db, 1)
	if err != nil {
		log.Println(err)
		return
	}

	// read the CSV file
	countries, err := ReadCSV("countries.csv")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("countries size: " + strconv.Itoa(len(countries)))
	// insert the data into the SQLite database
	for _, country := range countries {
		_, err := Insert(db, &country)
		if err != nil {
			log.Println(err)
			break
		}
	}

	country2, err := FindById(db, 1)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(country2.Name)

	// find all countries
	countries3, err := FindAll(db)
	if err != nil {
		log.Println(err)
		return
	}
	for _, c := range countries3 {
		log.Printf("%s\n", c.Name)
	}
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY,
        name     TEXT NOT NULL,
        population INTEGER NOT NULL,
        area INTEGER NOT NULL
    );`

	return db.Exec(sql)
}

func Insert(db *sql.DB, c *Country) (int64, error) {
	sql := `INSERT INTO countries (name, population, area) 
            VALUES (?, ?, ?);`

	// stmt,err:=db.Prepare(sql)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// _, err = stmt.Exec(c.Name, c.Population, c.Area)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result, err := db.Exec(sql, c.Name, c.Population, c.Area)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func Update(db *sql.DB, id int, population int) (int64, error) {
	sql := `UPDATE countries SET population = ? WHERE id = ?;`
	result, err := db.Exec(sql, population, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func Delete(db *sql.DB, id int) (int64, error) {
	sql := `DELETE FROM countries WHERE id = ?`

	// stmt,err:=db.Prepare(sql)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = stmt.Exec(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result, err := db.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func ReadCSV(filename string) ([]Country, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Parse the CSV file
	var countries []Country
	for _, record := range records[1:] { // Skip header row
		population, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		area, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		country := Country{
			Name:       record[0],
			Population: population,
			Area:       area,
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func FindById(db *sql.DB, id int) (*Country, error) {
	sql := `SELECT * FROM countries WHERE id = ?`
	row := db.QueryRow(sql, id)
	c := &Country{}
	err := row.Scan(&c.Id, &c.Name, &c.Population, &c.Area)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FindAll(db *sql.DB) ([]Country, error) {
	sql := `SELECT * FROM countries ORDER BY name`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		c := &Country{}
		err := rows.Scan(&c.Id, &c.Name, &c.Population, &c.Area)
		if err != nil {
			return nil, err
		}
		countries = append(countries, *c)
	}
	return countries, nil
}
