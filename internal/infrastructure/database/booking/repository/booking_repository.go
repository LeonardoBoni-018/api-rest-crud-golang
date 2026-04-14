package repository

import (
	"context"
	"os"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository/entity/converter"
)

const bookingCollectionEnv = "MONGODB_BOOKING_COLLECTION"

type bookingRepository struct {
	database *mongo.Database
}

func NewBookingRepository(database *mongo.Database) BookingRepository {
	return &bookingRepository{database: database}
}

type BookingRepository interface {
	CreateBooking(booking *domain.Booking) (*domain.Booking, *rest_err.RestErr)
	FindBookingsByTenantID(tenantID string) ([]*domain.Booking, *rest_err.RestErr)
	FindBookingByID(id string) (*domain.Booking, *rest_err.RestErr)
}

func (r *bookingRepository) collection() *mongo.Collection {
	collectionName := os.Getenv(bookingCollectionEnv)
	return r.database.Collection(collectionName)
}

func (r *bookingRepository) CreateBooking(b *domain.Booking) (*domain.Booking, *rest_err.RestErr) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	result, err := r.collection().InsertOne(context.Background(), converter.ToEntity(b))

	if err != nil {
		logger.Error("Error trying to insert booking", err, zap.String("journey", "createBooking"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	b.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return b, nil
}

func (r *bookingRepository) FindBookingsByTenantID(tenantID string) ([]*domain.Booking, *rest_err.RestErr) {
	filter := bson.D{{Key: "tenant_id", Value: tenantID}}
	cursor, err := r.collection().Find(context.Background(), filter)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer cursor.Close(context.Background())

	var bookings []*domain.Booking
	for cursor.Next(context.Background()) {
		var entity entity.BookingEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		bookings = append(bookings, converter.ToDomain(entity))
	}
	return bookings, nil
}

func (r *bookingRepository) FindBookingByID(id string) (*domain.Booking, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid booking ID")
	}

	doc := &entity.BookingEntity{}
	filter := bson.D{{Key: "_id", Value: objID}}
	if err := r.collection().FindOne(context.Background(), filter).Decode(doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Booking not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return converter.ToDomain(*doc), nil
}
