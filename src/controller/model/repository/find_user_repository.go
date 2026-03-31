package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity/converter"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserbyEmail repository", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmail repository executed successfully", zap.String("journey", "findUserByEmail"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init FindUserById repository", zap.String("journey", "FindUserById"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, zap.String("journey", "FindUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, zap.String("journey", "FindUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserById repository executed successfully", zap.String("journey", "FindUserById"), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init FindUserByEmailAndPassword repository", zap.String("journey", "FindUserByEmailAndPassword"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password invalid"
			logger.Error(errorMessage, zap.String("journey", "FindUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, zap.String("journey", "FindUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmailAndPassword repository executed successfully", zap.String("journey", "FindUserByEmailAndPassword"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
