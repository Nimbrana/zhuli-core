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

	Server     string
	Database   string
	Collection string

	User     string
	Password string
}

var mongoDriver = MongoDBDriver{}
var fsDriver = FileSystemDriver{"PATH"}

// Init integrated drivers
func (dc *DataController) Init() error {
	if dc.UseMongoDriver {
		return configureMongoDB(dc)
	} else if dc.UseFileSystemDriver {
		return errors.New("FileSystemDriver " + notImplemented)
	} else if dc.UseMongoDriver && dc.UseFileSystemDriver {
		return errors.New(cannotUseBoth)
	}

	return errors.New(notConfigured)
}

// configureMongoDB Initialize the connection with MongoDB
func configureMongoDB(dc *DataController) error {
	mongoDriver.Server = dc.Server
	mongoDriver.Database = dc.Database
	mongoDriver.Collection = dc.Collection

	return mongoDriver.Connect()
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

// Update an object of the Database or FileSystem (not implemented yet).
func (dc *DataController) Update(object interface{}) error {
	if dc.UseMongoDriver {
		return mongoDriver.Update(object)
	} else if dc.UseFileSystemDriver {
		return errors.New("FileSystemDriver " + notImplemented)
	}

	return errors.New(notConfigured)
}

// Delete an object of the Database or FileSystem (not implemented yet).
func (dc *DataController) Delete(object interface{}) error {
	if dc.UseMongoDriver {
		return mongoDriver.Delete(object)
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

// GetByID returns an object by ID
func (dc *DataController) GetByID(id string) (interface{}, error) {
	if dc.UseMongoDriver {
		return mongoDriver.FindByID(id)
	} else if dc.UseFileSystemDriver {
		return nil, errors.New("FileSystemDriver " + notImplemented)
	}

	return nil, errors.New(notConfigured)
}
