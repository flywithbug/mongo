package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var sessionMap = make(map[string]*tMongo)

//
type Monger interface {
	IsDestroyed() bool
	Destroy()
	Use(db string) error
	Session() *mgo.Session
	CurrentDB() string
	Collection(c string) *mgo.Collection

	// CURD
	Insert(collection string, docs ...interface{}) (interface{}, error)
	Update(collection string, selector interface{}, update interface{}) error
	FindOne(collection string, query interface{}, result interface{}) (interface{}, error)
	FindAll(collection string, query interface{}, results interface{}) ([]interface{}, error)
	Remove(collection string, selector interface{}) error
	RemoveAll(collection string, selector interface{}) error
}

func CloseAll()  {
	for _,v := range sessionMap{
		v.session.Close()
	}
}


type tMongo struct {
	db        string
	destroyed bool
	session   *mgo.Session
}

func (s tMongo) IsDestroyed() bool {
	return s.destroyed
}

func (s *tMongo) Destroy() {
	if s.destroyed {
		return
	}
	s.destroyed = true
	s.session.Close()
}

func (s *tMongo) Use(db string) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c, ok := sessionMap[db]
	if ok {
		s.db = c.db
		s.session = c.session.Clone()
	}
	return errors.New(ErrNoConnection + " named " + db)
}

func (s *tMongo) Session() *mgo.Session {
	if s.destroyed {
		return nil
	}
	return s.session
}

func (s *tMongo) CurrentDB() string {
	return s.db
}

func (s *tMongo) Collection(c string) *mgo.Collection {
	if s.destroyed {
		return nil
	}
	return s.session.DB(s.db).C(c)
}

