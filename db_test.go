package mongo

import (
	"fmt"
	"testing"
)

func TestIncrement(t *testing.T) {
	RegisterMongo("127.0.0.1:27017", "test")
	count := 0
	for {
		if count >= 1000 {
			break
		}
		count++
		inId, err := GetIncrementId("test", "userId")
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println("IncrementId:", inId)
	}
}
