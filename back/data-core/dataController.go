package datacore

import "errors"

const (
	notConfigured  = "DataController not configured"
	notImplemented = "not implemented yet"
	cannotUseBoth  = "Cannot use both drivers yet"
)

// DataController manage the data flow
type DataController struct {
	UseMongoDriver      bool
	UseFileSystemDriver bool
}

var mongoDriver = MongoDBDriver{"SERVER", "DATABASE", "COLLECTION"}
var fsDriver = FileSystemDriver{"PATH"}

// Init integrated drivers
func (dc *DataController) Init() error {
	if dc.UseMongoDriver {
		return mongoDriver.Connect()
	} else if dc.UseFileSystemDriver {
		return errors.New("FileSystemDriver " + notImplemented)
	} else if dc.UseMongoDriver && dc.UseFileSystemDriver {
		return errors.New(cannotUseBoth)
	}

	return errors.New(notConfigured)
}

// Save an object to Database or FileSystem (not implemented yet).
func (dc *DataController) Save(object interface{}) error {
	if dc.UseMongoDriver {
		return mongoDriver.Insert(object)
	} else if dc.UseFileSystemDriver {
		return errors.New("FileSystemDriver " + notImplemented)
	}

	return errors.New(notConfigured)
}

// GetAll returns all data found in the specified collection
func (dc *DataController) GetAll() ([]interface{}, error) {
	if dc.UseMongoDriver {
		return mongoDriver.FindAll()
	} else if dc.UseFileSystemDriver {
		return nil, errors.New("FileSystemDriver " + notImplemented)
	}

	return nil, errors.New(notConfigured)
}
