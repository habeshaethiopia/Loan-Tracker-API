package repositories

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type logRepository struct {
	database   mongo.Database
	collection string
}

// GetAllLogs implements domain.LogRepository.
func (l *logRepository) GetAllLogs(ctx context.Context) ([]domain.LogEntry, error) {
	collection := l.database.Collection(l.collection)
	var log []domain.LogEntry
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var logEntry domain.LogEntry
		err := cursor.Decode(&logEntry)
		if err != nil {
			return nil, err
		}
		log = append(log, logEntry)
	}
	return log, nil
}

// GetLogsByEventType implements domain.LogRepository.
func (l *logRepository) GetLogsByEventType(ctx context.Context, eventType string) ([]domain.LogEntry, error) {
	collection := l.database.Collection(l.collection)
	var log []domain.LogEntry
	cursor, err := collection.Find(ctx, bson.M{"event_type": eventType})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var logEntry domain.LogEntry
		err := cursor.Decode(&logEntry)
		if err != nil {
			return nil, err
		}
		log = append(log, logEntry)
	}
	return log, nil
}

// GetLogsByUserID implements domain.LogRepository.
func (l *logRepository) GetLogsByUserID(ctx context.Context, userID string) ([]domain.LogEntry, error) {
	collection := l.database.Collection(l.collection)
	var log []domain.LogEntry
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var logEntry domain.LogEntry
		err := cursor.Decode(&logEntry)
		if err != nil {
			return nil, err
		}
		log = append(log, logEntry)
	}
	return log, nil

}

// SaveLog implements domain.LogRepository.
func (l *logRepository) SaveLog(ctx context.Context, log domain.LogEntry) error {
	collection := l.database.Collection(l.collection)
	_, err := collection.InsertOne(ctx, log)
	return err
}

func NewLogRepository(database mongo.Database, collection string) domain.LogRepository {
	return &logRepository{
		database:   database,
		collection: collection,
	}
}
