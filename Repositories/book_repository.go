package repositories

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	database   mongo.Database
	collection string
}

// CreateBook implements domain.BookRepository.
func (b *BookRepository) CreateBook(ctx context.Context, book *domain.Book) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.InsertOne(ctx, book)
	return err
}

// DeleteBook implements domain.BookRepository.
func (b *BookRepository) DeleteBook(ctx context.Context, id primitive.ObjectID) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err

}

// GetAllAvailableBooks implements domain.BookRepository.
func (b *BookRepository) GetAllAvailableBooks(ctx context.Context) ([]*domain.Book, error) {
	collection := b.database.Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{"is_available": true})
	if err != nil {
		return nil, err
	}
	var books []*domain.Book
	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, nil
}

// GetAllBooks implements domain.BookRepository.
func (b *BookRepository) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	collection := b.database.Collection(b.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var books []*domain.Book
	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByID implements domain.BookRepository.
func (b *BookRepository) GetBookByID(ctx context.Context, id primitive.ObjectID) (*domain.Book, error) {
	collection := b.database.Collection(b.collection)
	var book domain.Book
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// UpdateBook implements domain.BookRepository.
func (b *BookRepository) UpdateBook(ctx context.Context, book *domain.Book) error {
	collection := b.database.Collection(b.collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": book.ID}, bson.M{"$set": book})
	return err
}

func NewBookRepository(db mongo.Database, collection string) domain.BookRepository {
	return &BookRepository{
		database:   db,
		collection: collection,
	}
}
