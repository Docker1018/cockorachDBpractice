package cockroachdb

import (
	"cockroach/model"
	"database/sql"
	"log"
)

func PrintBalances(db *sql.DB) []*model.Member {
	rows, err := db.Query("SELECT * FROM member")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	vals := make([]*model.Member, 0)
	for rows.Next() {
		var Account, Password, Name string
		if err := rows.Scan(&Account, &Password, &Name); err != nil {
			log.Fatal(err)
		}
		vals = append(vals, &model.Member{
			Account: Account,
			Password: Password,
			Name: Name,
		})
	}
	return vals
}

func insert(db *sql.DB,) error{
	// Insert two rows into the "accounts" table.

	stmt, err := db.Prepare("INSERT INTO member (account, password, name) VALUES ($1, $2, $3, $4)")
	if _, err := db.Exec(
	    "INSERT INTO member (id, balance) VALUES (3, 1000), (2, 250)") 
			err != nil {
	    return err
	}
}
