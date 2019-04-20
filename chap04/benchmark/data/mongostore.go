package data

import "gopkg.in/mgo.v2"

type MongoStore struct {
	session *mgo.Session
}

func NewMongoStore(connect string) (*MongoStore, error) {
	session, err := mgo.Dial(connect)
	if err != nil {
		return nil, err
	}

	return &MongoStore{session: session}, nil
}

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

func (m *MongoStore) DeleteAllKittens() {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").DropCollection()
}

func (m *MongoStore) InsertKittens(kittens []Kitten) {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").Insert(kittens)
}
