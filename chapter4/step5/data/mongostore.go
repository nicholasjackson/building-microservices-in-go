package data

import "labix.org/v2/mgo"

// MongoStore is a MongoDB data store which implements the Store interface
type MongoStore struct {
	session *mgo.Session
}

// NewMongoStore creates an instance of MongoStore with the given connection string
func NewMongoStore(connection string) (*MongoStore, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	return &MongoStore{session: session}, nil
}

// Search returns Kittens from the MongoDB instance which have the name name
func (m *MongoStore) Search(name string) []Kitten {
	s := m.session.Clone()
	defer s.Close()

	var results []Kitten
	c := s.DB("kittenserver").C("kittens")
	err := c.Find(Kitten{Name: name}).All(&results)
	if err != nil {
		return nil
	}

	return results
}

// DeleteAllKittens deletes all the kittens from the datastore
func (m *MongoStore) DeleteAllKittens() {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").DropCollection()
}

// InsertKittens inserts a slice of kittens into the datastore
func (m *MongoStore) InsertKittens(kittens []Kitten) {

	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").Insert(kittens)
}
