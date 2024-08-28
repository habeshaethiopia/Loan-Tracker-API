package repositories

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookBorrowRepository struct {
	database   mongo.Database
	collection string
}

// CreateBorrowRequest implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) CreateBorrowRequest(ctx context.Context, request *domain.BorrowRequest) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.InsertOne(ctx, request)
	return err

}

// DeleteBorrowRequest implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) DeleteBorrowRequest(ctx context.Context, id primitive.ObjectID) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// GetAllBorrowRequests implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) GetAllBorrowRequests(ctx context.Context) ([]*domain.BorrowRequest, error) {
	collection := b.database.Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var requests []*domain.BorrowRequest
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, err
	}
	return requests, nil

}

// GetBorrowRequestByID implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) GetBorrowRequestByID(ctx context.Context, id primitive.ObjectID) (*domain.BorrowRequest, error) {
	collection := b.database.Collection(b.collection)
	var request domain.BorrowRequest
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// GetBorrowRequestsByBookID implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) GetBorrowRequestsByBookID(ctx context.Context, bookID primitive.ObjectID) ([]*domain.BorrowRequest, error) {
	collection := b.database.Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{"book_id": bookID})
	if err != nil {
		return nil, err
	}
	var requests []*domain.BorrowRequest
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, err
	}
	return requests, nil

}

// GetBorrowRequestsByUserID implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) GetBorrowRequestsByUserID(ctx context.Context, userID primitive.ObjectID) ([]*domain.BorrowRequest, error) {
	collection := b.database.Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	var requests []*domain.BorrowRequest
	if err = cursor.All(ctx, &requests); err != nil {
		return nil, err
	}
	return requests, nil
}

// UpdateBorrowRequest implements domain.BorrowRequestRepository.
func (b *BookBorrowRepository) UpdateBorrowRequest(ctx context.Context, request *domain.BorrowRequest) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": request.ID}, bson.M{"$set": request})
	return err
}

func NewBookBorrowRepository(database mongo.Database, collection string) domain.BorrowRequestRepository {
	return &BookBorrowRepository{
		database:   database,
		collection: collection,
	}
}
