package mongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/PatrickKvartsaniy/image-processing-service/errors"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Config struct {
		URI        string
		Database   string
		Collection string
	}

	Repository struct {
		client     *mongo.Client
		collection *mongo.Collection
	}
)

func New(ctx context.Context, cfg Config) (*Repository, error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Repository{
		client:     client,
		collection: client.Database(cfg.Database).Collection(cfg.Collection),
	}, nil
}

func (r Repository) GetImage(ctx context.Context, id string) (*model.Image, error) {
	res := r.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if res.Err() != nil {
		return nil, processError(res.Err())
	}

	var i model.Image
	if err := res.Decode(&i); err != nil {
		return nil, processError(err)
	}

	return &i, nil
}

func (r Repository) GetMultipleImages(ctx context.Context, limit, offset int64) ([]*model.Image, error) {
	cur, err := r.collection.Find(ctx, bson.D{}, findOptions(limit, offset))
	if err != nil {
		return nil, processError(err)
	}

	var images []*model.Image
	for cur.Next(ctx) {
		var img model.Image
		if err := cur.Decode(&img); err != nil {
			return nil, processError(err)
		}
		images = append(images, &img)
	}

	if err := cur.Err(); err != nil {
		return nil, processError(err)
	}

	if err := cur.Close(ctx); err != nil {
		return nil, processError(err)
	}

	return images, nil
}

func (r Repository) SaveImage(ctx context.Context, image *model.Image) error {
	if image == nil {
		return errors.InvalidInput
	}
	_, err := r.collection.InsertOne(ctx, *image)
	return processError(err)
}

func (r Repository) UpdateImage(ctx context.Context, image *model.Image) error {
	if image == nil {
		return errors.InvalidInput
	}

	filter := bson.D{
		{Key: "_id", Value: image.ID},
		{Key: "version", Value: image.Version},
	}

	_, err := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: *image}}, options.Update())
	if err != nil {
		return processError(err)
	}

	return nil
}

func (r Repository) Close() {
	if err := r.client.Disconnect(context.TODO()); err != nil {
		logrus.Error(err)
	}
}

func (r Repository) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return r.client.Ping(ctx, nil)
}

func findOptions(limit, offset int64) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)
	return findOptions
}

func processError(err error) error {
	if err == nil {
		return nil
	}
	if err == mongo.ErrNoDocuments {
		return errors.NotFound
	}

	logrus.Error(err)

	return errors.Internal
}
