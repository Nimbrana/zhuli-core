package datacore

import "errors"

// DataController manage the data flow
type DataController struct {
	UseMongoDriver      bool
	UseFileSystemDriver bool
}

var mongoDriver = MongoDBDriver{"SERVER", "DATABASE", "COLLECTION"}
var fsDriver = FileSystemDriver{"PATH"}

// Save an object to Database or FileSystem (not implemented yet).
func (dc *DataController) Save(object interface{}) error {

	err := checkConfiguration(dc)

	if err != nil {
		return err
	}

	if dc.UseMongoDriver {
		return mongoDriver.Insert(object)
	}

	return err
}

// GetAll returns all data found in the specified collection
func (dc *DataController) GetAll() ([]interface{}, error) {
	err := checkConfiguration(dc)

	if err != nil {
		return nil, err
	}

	if dc.UseMongoDriver {
		return mongoDriver.FindAll()
	}

	return nil, err
}

func checkConfiguration(dc *DataController) error {
	if dc.UseMongoDriver && dc.UseFileSystemDriver {
		return errors.New("Cannot use both drivers yet")
	} else if dc.UseMongoDriver {
		if mongoDriver.Ping() != nil {
			return mongoDriver.Connect()
		}
		return nil
	} else if dc.UseFileSystemDriver {
		return errors.New("FileSystemDriver not implemented yet")
	}

	return errors.New("DataController not configured")
}
