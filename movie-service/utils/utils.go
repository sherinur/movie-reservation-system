package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
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

func GenerateID() string {
	id := uuid.New()

	ID := base64.URLEncoding.EncodeToString(id[:])

	return ID
}
