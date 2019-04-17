package data

import "gopkg.in/mgo.v2"

// Store 인터페이스를 구현 한 몽고 디비
type MongoStore struct {
	session *mgo.Session
}

// 주어진 connection string 을 가지고 MongoStore 인스턴스 생성
func NewMongoStore(connection string) (*MongoStore, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	return &MongoStore{session: session}, nil
}

// kitten name 으로 검색
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

// 전체 kitten 삭제
func (m *MongoStore) DeleteAllKittens() {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").DropCollection()
}

func (m *MongoStore) InsertKittens(kitten []Kitten) {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").Insert(kitten)
}
