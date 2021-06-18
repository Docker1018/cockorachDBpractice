package controller

import (
	"cockroach/cockroachdb"
	"log"
)
func SelectData () {
	db, err := cockroachdb.Connect()
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	cockroachdb.PrintBalances(db);
}
func InsertData () {}
func UpdateData () {}
func DeleteData () {}