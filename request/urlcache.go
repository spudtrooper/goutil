package request

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ()

type urlCache struct {
	db *mongo.Database
}

type decodableResponse struct {
	Data    []byte
	Cookies []Cookie
}

type storedResponse struct {
	URI      string
	Response decodableResponse
}

func MakeURLCacheFromDB(db *mongo.Database) *urlCache {
	return &urlCache{
		db: db,
	}
}

func ConnectToURLCache(ctx context.Context, cOpts ...ConnectToURLCacheOption) (*urlCache, error) {
	opts := MakeConnectToURLCacheOptions(cOpts...)

	port := or.Int(opts.Port(), 27017)
	dbName := or.String(opts.DbName(), "goutil")

	uri := fmt.Sprintf("mongodb://localhost:%d", port)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Errorf("mongo.Connect: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Errorf("client.Ping: %v", err)
	}

	log.Printf("connected to %q @ %s", dbName, uri)

	db := client.Database(dbName)
	res := MakeURLCacheFromDB(db)
	return res, nil
}

func (c *urlCache) SaveRequest(ctx context.Context, uri string, res Response) error {
	filter := bson.D{{Key: "uri", Value: uri}}
	if _, err := c.db.Collection("requests").DeleteMany(ctx, filter); err != nil {
		return errors.Errorf("DeleteMany: %v", err)
	}
	s := storedResponse{
		URI: uri,
		Response: decodableResponse{
			Data:    res.Data,
			Cookies: res.Cookies,
		},
	}
	if _, err := c.db.Collection("requests").InsertOne(ctx, s); err != nil {
		return err
	}
	return nil
}

func isEmptyResultError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "no documents in result")
}

func (c *urlCache) FindRequest(ctx context.Context, uri string) (*Response, error) {
	filter := bson.D{{Key: "uri", Value: uri}}
	resp := c.db.Collection("requests").FindOne(ctx, filter)
	if resp.Err() != nil {
		if !isEmptyResultError(resp.Err()) {
			return nil, resp.Err()
		}
		return nil, nil
	}
	var s storedResponse
	if err := resp.Decode(&s); err != nil {
		return nil, err
	}
	res := &Response{
		Data:    s.Response.Data,
		Cookies: s.Response.Cookies,
	}
	return res, nil
}
