package mongo

func DropDatabase(name string)error  {
	_,db := getDb(name)
	if db != nil {
	  return 	db.DropDatabase()
	}
	return nil
}

func DropCollection(db,collection string)error  {
	_,c :=connect(db,collection)
	if c != nil {
		return c.DropCollection()
	}
	return nil
}
