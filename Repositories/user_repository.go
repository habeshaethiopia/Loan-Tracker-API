package repositories

import (
	"context"
	"fmt"

	domain "LoanTrackerAPI/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	var user1 domain.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user1)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	_, err = collection.InsertOne(ctx, user)

	return err
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	
	collection := u.database.Collection(u.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": idHex})
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers implements domain.UserRepository.
func (u *userRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var users []*domain.User
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user *domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByEmail implements domain.UserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID implements domain.UserRepository.
func (ur *userRepository) GetUserByID(c context.Context, id string) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.User{}, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)

	if err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)

	updateResult, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v documents in the trainers collection\n", updateResult.ModifiedCount)
	return nil
}

// VerifyEmail implements domain.UserRepository.
func (u *userRepository) VerifyEmail(ctx context.Context, token string, email string) error {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return err
	}
	if user.IsVerified {
		return fmt.Errorf("user with email %s is already verified", email)
	}

	err = collection.FindOne(ctx, bson.M{"email": email, "verify_token": token}).Decode(&user)
	if err != nil {
		return err
	}
	user.IsVerified = true
	user.VerifyToken = ""

	res, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v documents in the userd collection\n", res)
	return nil
}

// GetUserByEmail implements domain.UserRepository.

// GetTaskByEmail implements domain.UserRepository.

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
