package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Convert struct to bson.D format
func ConvertToBsonD(movie interface{}) (interface{}, error) {
	if movie == nil {
		return nil, fmt.Errorf("input movie is nil")
	}

	bsonData, err := bson.Marshal(movie)
	if err != nil {
		return nil, fmt.Errorf("error marshaling movie: %w", err)
	}

	var bsonDoc bson.D
	if err := bson.Unmarshal(bsonData, &bsonDoc); err != nil {
		return nil, fmt.Errorf("error unmarshaling BSON: %w", err)
	}

	return bsonDoc, nil
}

// func IsEmpty(v interface{}) string {
// 	if v == nil {
// 		return "Nil"
// 	}

// 	val := reflect.ValueOf(v)
// 	if val.Kind() == reflect.Ptr && val.IsNil() {
// 		return "Nil pointer"
// 	}

// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}

// 	switch val.Kind() {
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			field := val.Field(i)
// 			fieldName := val.Type().Field(i).Name
// 			if IsEmpty(field.Interface()) != "" {
// 				return fieldName
// 			}
// 		}
// 	case reflect.Slice, reflect.Array:
// 		for i := 0; i < val.Len(); i++ {
// 			if IsEmpty(val.Index(i).Interface()) != "" {
// 				return fmt.Sprintf("[%d]", i)
// 			}
// 		}
// 	case reflect.String:
// 		if val.String() == "" {
// 			return "String"
// 		}
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		if val.Int() == 0 {
// 			return "Integer"
// 		}
// 	case reflect.Float32, reflect.Float64:
// 		if val.Float() == 0.0 {
// 			return "Float"
// 		}
// 	}

// 	return ""
// }
