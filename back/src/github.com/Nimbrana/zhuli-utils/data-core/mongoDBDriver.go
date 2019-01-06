package datacore

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDBDriver represent an interface to mongodb including de Server url, Database and Collection names.
type MongoDBDriver struct {
	Server     string
	Database   string
	Collection string
}

var mDatabase *mgo.Database

// Connect to MongoDB Server
func (d *MongoDBDriver) Connect() {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Print(err)
	}
	mDatabase = session.DB(d.Database)
}

// FindAll find all documents into collection and return an array of Bson maps
func (d *MongoDBDriver) FindAll() ([]bson.M, error) {
	var result []bson.M
	err := mDatabase.C(d.Collection).Find(bson.M{}).All(&result)
	return result, err
}

// FindByID find a Bson document by id
func (d *MongoDBDriver) FindByID(id string) ([]bson.M, error) {
	var result []bson.M
	err := mDatabase.C(d.Collection).FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

// Insert a Bson document into mongoDB colletion
func (d *MongoDBDriver) Insert(object bson.M) error {
	err := mDatabase.C(d.Collection).Insert(&object)
	return err
}

// Delete an existing Bson document
func (d *MongoDBDriver) Delete(object bson.M) error {
	err := mDatabase.C(d.Collection).Remove(&object)
	return err
}

// Update an existing Bson document
func (d *MongoDBDriver) Update(object bson.M) error {
	err := mDatabase.C(d.Collection).UpdateId(object["_id"], &object)
	return err
}
