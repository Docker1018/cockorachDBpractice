package controller

import (
	"cockroach/cockroachdb"
	"cockroach/model"
	"log"
)

func SelectData() []*model.Member {
	db, err := cockroachdb.Connect()
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	return cockroachdb.PrintBalances(db)
}
func InsertData(d *model.Member) *model.Member {
	db, err ;= cockroachdb.Connect()
	
	return d
}
func UpdateData() {}
func DeleteData() {}
