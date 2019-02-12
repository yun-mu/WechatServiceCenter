package model

import "model/db"

func updateDoc(tableName string, query, update interface{}) error {
	cntrl := db.NewCloneMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(tableName)
	return table.Update(query, update)
}

func insertDocs(tableName string, docs ...interface{}) error {
	cntrl := db.NewCloneMgoDBCntlr()
	defer cntrl.Close()
	table := cntrl.GetTable(tableName)
	return table.Insert(docs...)
}
