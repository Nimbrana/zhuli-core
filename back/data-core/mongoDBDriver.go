package datacore

import (
	"errors"

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
func (d *MongoDBDriver) Connect() error {
	session, err := mgo.Dial(d.Server)
	if err == nil {
		mDatabase = session.DB(d.Database)
	}

	return err
}

// FindAll find all documents into collection and return an array of Bson maps
func (d *MongoDBDriver) FindAll() ([]interface{}, error) {
	var result []interface{}
	err := mDatabase.C(d.Collection).Find(bson.M{}).All(&result)
	return result, err
}

// FindByID find a Bson document by id
func (d *MongoDBDriver) FindByID(id string) (interface{}, error) {
	var result interface{}
	err := mDatabase.C(d.Collection).FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

// Insert a Bson document into mongoDB colletion
func (d *MongoDBDriver) Insert(object interface{}) error {
	data, err := unmarshal(object)
	if err != nil {
		return err
	}
	err = mDatabase.C(d.Collection).Insert(&data)
	return err
}

// Delete an existing Bson document
func (d *MongoDBDriver) Delete(object interface{}) error {
	err := mDatabase.C(d.Collection).Remove(&object)
	return err
}

// Update an existing Bson document
func (d *MongoDBDriver) Update(id string, object interface{}) error {
	err := mDatabase.C(d.Collection).UpdateId(id, &object)
	return err
}

// Ping is that a simple ping
func (d *MongoDBDriver) Ping() error {
	return mDatabase.Session.Ping()
}

// unmarshal converts an interface into a map of values or return an error if exist an assert in the parsing process
func unmarshal(object interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	binaryData, ok := object.([]byte)
	bson.Unmarshal(binaryData, &data)

	var err error

	if !ok {
		err = errors.New("Parsing error")
	}
	return data, err
}
