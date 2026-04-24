package repository

import (
	"context"
	"os"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/plan"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/payment/repository/entity/converter"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/payment/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const subscriptionCollectionEnv = "MONGODB_SUBSCRIPTION_COLLECTION"
const paymentCollectionEnv = "MONGODB_PAYMENT_COLLECTION"

type paymentRepository struct {
	database *mongo.Database
}

func NewPaymentRepository(database *mongo.Database) *paymentRepository {
	return &paymentRepository{database: database}
}

type PaymentRepository interface {
	CreateSubscription(sub *plan.Subscription) (*plan.Subscription, *rest_err.RestErr)
	FindSubscriptionById(id string) (*plan.Subscription, *rest_err.RestErr)
	FindSubscriptionByTenantId(tenantID string) (*plan.Subscription, *rest_err.RestErr)
	UpdateSubscription(id string, sub *plan.Subscription) (*plan.Subscription, *rest_err.RestErr)
	CreatePayment(payment *plan.Payment) (*plan.Payment, *rest_err.RestErr)
	FindPaymentsByTenantId(tenantID string) ([]plan.Payment, *rest_err.RestErr)
}

func (r *paymentRepository) subscriptionCollection() *mongo.Collection {
	collectionName := os.Getenv(subscriptionCollectionEnv)
	if collectionName == "" {
		collectionName = "subscriptions"
	}
	return r.database.Collection(collectionName)
}

func (r *paymentRepository) paymentCollection() *mongo.Collection {
	collectionName := os.Getenv(paymentCollectionEnv)
	if collectionName == "" {
		collectionName = "payments"
	}
	return r.database.Collection(collectionName)
}

func (r *paymentRepository) CreateSubscription(sub *plan.Subscription) (*plan.Subscription, *rest_err.RestErr) {
	sub.CreatedAt = time.Now()
	sub.UpdatedAt = time.Now()

	result, err := r.subscriptionCollection().InsertOne(context.Background(), converter.ToEntitySubscription(sub))
	if err != nil {
		return nil, rest_err.NewInternalServerError("error creating subscription")
	}

	sub.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return sub, nil
}

func (r *paymentRepository) FindSubscriptionById(id string) (*plan.Subscription, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid subscription ID")
	}

	doc := &entity.SubscriptionEntity{}
	err = r.subscriptionCollection().FindOne(context.Background(), bson.M{"_id": objID}).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("subscription not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ToDomainSubscription(doc), nil
}

func (r *paymentRepository) FindSubscriptionByTenantId(tenantID string) (*plan.Subscription, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(tenantID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid tenant ID")
	}

	doc := &entity.SubscriptionEntity{}
	err = r.subscriptionCollection().FindOne(context.Background(), bson.M{"tenant_id": objID}).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("subscription not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ToDomainSubscription(doc), nil
}

func (r *paymentRepository) UpdateSubscription(id string, sub *plan.Subscription) (*plan.Subscription, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid subscription ID")
	}

	sub.UpdatedAt = time.Now()
	_, err = r.subscriptionCollection().UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": converter.ToEntitySubscription(sub)},
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return sub, nil
}

func (r *paymentRepository) CreatePayment(p *plan.Payment) (*plan.Payment, *rest_err.RestErr) {
	p.CreatedAt = time.Now()

	result, err := r.paymentCollection().InsertOne(context.Background(), converter.ToEntityPayment(p))
	if err != nil {
		return nil, rest_err.NewInternalServerError("error creating payment")
	}

	p.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return p, nil
}

func (r *paymentRepository) FindPaymentsByTenantId(tenantID string) ([]plan.Payment, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(tenantID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid tenant ID")
	}

	cursor, err := r.paymentCollection().Find(context.Background(), bson.M{"tenant_id": objID})
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer cursor.Close(context.Background())

	var payments []plan.Payment
	for cursor.Next(context.Background()) {
		var doc entity.PaymentEntity
		if err := cursor.Decode(&doc); err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		payments = append(payments, *converter.ToDomainPayment(&doc))
	}

	return payments, nil
}