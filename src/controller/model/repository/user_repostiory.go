package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"

)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	// UpdateUser(
	// 	userId string,
	// 	userDomain model.UserDomainInterface,
	// ) *rest_err.RestErr

	// DeleteUser(
	// 	userId string,
	// ) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	// FindUserByEmailAndPassword(
	// 	email string,
	// 	password string,
	// ) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
