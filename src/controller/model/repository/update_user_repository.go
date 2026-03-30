package repository

import (
	"context"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity/converter"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func (ur *userRepository) UpdateUser(
	userId string,

	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository", zap.String("journey", "updateUser"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User Updated successfully", zap.String("userId", userId), zap.String("journey", "updateUser"))
	return nil
}
