package datacore

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var dc = DataController{}
var movie = Movie{}

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

func TestInitDataControllerWithMongoDB(t *testing.T) {
	dc.UseMongoDriver = true
	dc.UseFileSystemDriver = false
	dc.Server = "localhost"
	dc.Database = "movies"
	dc.Collection = "movies"
	dc.User = ""
	dc.Password = ""

	if err := dc.Init(); err != nil {
		t.Errorf("%s", err)
	}
}

func TestInsertData(t *testing.T) {

	dc.UseMongoDriver = true
	dc.UseFileSystemDriver = false
	dc.Server = "localhost"
	dc.Database = "movies"
	dc.Collection = "movies"
	dc.User = ""
	dc.Password = ""

	movie.ID = bson.NewObjectId()
	movie.Name = "Dunkirk"
	movie.Description = "World war 2 movie"
	movie.CoverImage = "https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg"

	bMovie, _ := bson.Marshal(movie)

	if err := dc.Init(); err != nil {
		t.Errorf("%s", err)
	}

	if err := dc.Save(bMovie); err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetAllData(t *testing.T) {

	dc.UseMongoDriver = true
	dc.UseFileSystemDriver = false
	dc.Server = "localhost"
	dc.Database = "movies"
	dc.Collection = "movies"
	dc.User = ""
	dc.Password = ""

	// movie.ID = bson.NewObjectId()
	// movie.Name = "Dunkirk"
	// movie.Description = "World war 2 movie"
	// movie.CoverImage = "https://image.tmdb.org/t/p/w640/cUqEgoP6kj8ykfNjJx3Tl5zHCcN.jpg"

	//bMovie, _ := bson.Marshal(movie)

	if err := dc.Init(); err != nil {
		t.Errorf("%s", err)
	}

	data, err := dc.GetAll()

	t.Log(data)

	if err != nil {
		t.Errorf("%s", err)
	}
}
