package repository

import (
	"context"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(
	userId string,

) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "deleteUser"))

	collection_name := os.Getenv("MONGODB_USER_COLLECTION")
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User Deleted successfully", zap.String("userId", userId), zap.String("journey", "deleteUser"))
	return nil
}

