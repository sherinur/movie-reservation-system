package utils

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// Convert struct to bson.D format
func ConvertToBsonD(movie interface{}) (interface{}, error) {
	bsonData, err := bson.Marshal(movie)
	if err != nil {
		return nil, err
	}

	var bsonDoc bson.D
	if err := bson.Unmarshal(bsonData, &bsonDoc); err != nil {
		return nil, err
	}

	return bsonDoc, nil
}

func IsEmpty(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.IsValid() && val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
				return false
			}
		}
	}

	return true
}
