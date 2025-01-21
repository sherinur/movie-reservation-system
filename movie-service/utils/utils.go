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
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if IsEmpty(val.Field(i).Interface()) {
				return true
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if IsEmpty(val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0.0
	}

	return false
}
