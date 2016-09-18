package db

import (
	ara "github.com/diegogub/aranGO"
)

var DB *ara.Database
var NodeColl *ara.Collection
var RelColl *ara.Collection

var RelCollName = "tRel"
var NodeColName = "tNode"

func Connect() {
	session,err := ara.Connect("http://localhost:8529", "", "",false)
	if (err != nil) {
		panic(err)
	} else {
		DB = session.DB("test")
		if !DB.ColExist(RelCollName) {
			rel := ara.NewCollectionOptions(RelCollName, true)
			rel.IsEdge()
			DB.CreateCollection(rel)
		}
		NodeColl = DB.Col(NodeColName)
		RelColl = DB.Col(RelCollName)
	}
}

type Rel struct {
	ara.Edge
	Created int64
}

type Node struct {
	ara.Document
	Created int64
	Id int
}