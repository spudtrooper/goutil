package request

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createDBForTesting(ctx context.Context) (*mongo.Database, error) {
	const port = 27017
	const dbName = "goutiltest"

	uri := fmt.Sprintf("mongodb://localhost:%d", port)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return db, nil
}

func dropCacheForTesting(ctx context.Context) error {
	db, err := createDBForTesting(ctx)
	if err != nil {
		return err
	}
	if err := db.Drop(ctx); err != nil {
		return errors.Errorf("drop db: %+v", err)
	}
	return nil
}

func testFindRequest(ctx context.Context, t *testing.T, cache *urlCache) {
	uri := "http://foo.com/bar"
	resp, err := cache.FindRequest(ctx, uri)
	if err != nil {
		t.Fatalf("FindRequest: %+v", err)
	}
	if resp != nil {
		t.Errorf("FindRequest: expecting nil but got: %v", resp)
	}

	if err := cache.SaveRequest(ctx, uri, Response{Data: []byte("contents")}); err != nil {
		t.Fatalf("SaveRequest: %+v", err)
	}
	{
		resp, err := cache.FindRequest(ctx, uri)
		if err != nil {
			t.Fatalf("FindRequest: %+v", err)
		}
		if want, got := (&Response{Data: []byte("contents")}), resp; !reflect.DeepEqual(want, got) {
			t.Errorf("want <<< %+v >>> != got <<< %+v >>>", want, got)
		}
	}
}

func TestConnectToURLCache(t *testing.T) {
	ctx := context.Background()
	dropCacheForTesting(ctx)
	cache, err := ConnectToURLCache(ctx, ConnectToURLCacheDbName("goutiltest"))
	if err != nil {
		t.Fatalf("ConnectToURLCache: %+v", err)
	}
	testFindRequest(ctx, t, cache)
}

func TestFindQuery(t *testing.T) {
	ctx := context.Background()
	dropCacheForTesting(ctx)
	db, err := createDBForTesting(ctx)
	if err != nil {
		t.Fatalf("createDBForTesting: %+v", err)
	}
	cache := MakeURLCacheFromDB(db)
	testFindRequest(ctx, t, cache)
}
