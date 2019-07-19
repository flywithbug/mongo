package mongo

import(
	"gopkg.in/mgo.v2"
	"fmt"
)

type Index struct {
	Collection  string
	DBName  string
	Index  mgo.Index
	DropIndex  []string
}

func MgoIndex(list []Index,forceSync bool)error  {
	for _, v := range list{
		_,c := Collection(v.DBName,v.Collection)
		if len(v.DropIndex) > 0 {
			for _, dv := range v.DropIndex {
				if err := c.DropIndexName(dv);err != nil {
					return err
				}
			}
		}
		if forceSync {
			if err := c.DropIndexName(v.Index.Name);err != nil {
				fmt.Println(err)
			}
		}
		if err := c.EnsureIndex(v.Index);err != nil {
			return err
		}
	}
	return nil
}
