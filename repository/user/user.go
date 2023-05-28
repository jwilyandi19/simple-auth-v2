package user

import (
	"context"
	"log"
	"time"

	domain "github.com/jwilyandi19/simple-auth-v2/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, user *domain.User, id string) (*domain.User, error)
	Login(ctx context.Context, username string) (*domain.User, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	collection := u.db.Collection("users")

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("[UserRepository-GetByID] Error when getting primitive ID", err)
		return nil, err
	}

	filter := bson.M{"_id": idHex}
	var user domain.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println("[UserRepository-GetByID] Error when finding data", err)
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	collection := u.db.Collection("users")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("[UserRepository-GetAll] Error when finding data", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var users []domain.User
	for cur.Next(ctx) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			log.Println("[UserRepository-GetAll] Error when decoding data", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Println("[UserRepository-GetAll] Error from mongo", err)
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	collection := u.db.Collection("users")

	id := primitive.NewObjectID()
	user.ID = id
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("[UserRepository-Create] Error when inserting", err)
		return user, err
	}
	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User, id string) (*domain.User, error) {
	collection := u.db.Collection("users")
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("[UserRepository-Update] Error when getting primitive ID", err)
		return nil, err
	}

	filter := bson.M{"_id": idHex}

	update := bson.M{"$set": bson.M{
		"name":       user.Name,
		"username":   user.Username,
		"password":   user.Password,
		"is_admin":   user.IsAdmin,
		"updated_at": time.Now(),
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("[UserRepository-Update] Error when updating", err)
		return user, err
	}

	return user, nil
}

func (u *userRepository) Login(ctx context.Context, username string) (*domain.User, error) {
	collection := u.db.Collection("users")

	filter := bson.M{
		"username": username,
	}

	var user domain.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println("[UserRepository-Login] Wrong username + password with err", err)
		return nil, err
	}

	return &user, nil

}
