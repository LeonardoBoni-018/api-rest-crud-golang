package repository

import (
	"context"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity/converter"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "createUser"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), userEntity)
	if err != nil {
		logger.Error("Error trying to insert user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(primitive.ObjectID).Hex())

	return userDomain, nil
}
