package dbcopy

import (
	"context"
	"errors"
	"fmt"

	"github.com/skyrocket-qy/gox"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CopyMongoIndexes(ctx context.Context, srcCol, dstCol *mongo.Collection) (err error) {
	// Get indexes from source
	cursor, err := srcCol.Indexes().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list indexes: %w", err)
	}

	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for cursor.Next(ctx) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			return fmt.Errorf("failed to decode index: %w", err)
		}

		// Skip default _id index
		if name, ok := index["name"].(string); ok && name == "_id_" {
			continue
		}

		keys, ok := index["key"].(bson.M)
		if !ok {
			return fmt.Errorf(
				"index 'key' is not of type bson.M or does not exist: %v",
				index["key"],
			)
		}

		opts := options.Index()
		if name, ok := index["name"].(string); ok {
			opts.SetName(name)
		}

		if unique, ok := index["unique"].(bool); ok {
			opts.SetUnique(unique)
		}

		if sparse, ok := index["sparse"].(bool); ok {
			opts.SetSparse(sparse)
		}

		if expireAfterSeconds, ok := index["expireAfterSeconds"].(int32); ok {
			opts.SetExpireAfterSeconds(expireAfterSeconds)
		}

		model := mongo.IndexModel{
			Keys:    keys,
			Options: opts,
		}

		if _, err := dstCol.Indexes().CreateOne(ctx, model); err != nil {
			return fmt.Errorf("failed to create index: %w", err)
		}
	}

	return nil
}

func CopyMongoData(
	ctx context.Context,
	src, dst *mongo.Collection,
	filter bson.D,
	mode Mode,
	backupColName string,
) (err error) {
	cursor, err := src.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to find source documents: %w", err)
	}

	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var docs []any

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return fmt.Errorf("failed to decode document: %w", err)
		}

		delete(doc, "_id")
		docs = append(docs, doc)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %w", err)
	}

	if len(docs) == 0 {
		return errors.New("no documents to copy")
	}

	if backupColName != "" {
		if err := BackupMongoData(ctx, dst, filter, backupColName); err != nil {
			return fmt.Errorf("backup failed: %w", err)
		}
	}

	sess, err := dst.Database().Client().StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer sess.EndSession(ctx)

	callback := func(ctx mongo.SessionContext) error {
		switch mode {
		case ModeBasic:
			count, err := dst.CountDocuments(ctx, filter)
			if err != nil {
				return fmt.Errorf("failed to count destination documents: %w", err)
			}

			if count > 0 {
				return errors.New("destination already has data")
			}

		case ModeReplace:
			if _, err := dst.DeleteMany(ctx, filter); err != nil {
				return fmt.Errorf("failed to delete from destination: %w", err)
			}

		case ModeAppend:
			// Do nothing

		default:
			return fmt.Errorf("unsupported copy mode: %v", mode)
		}

		for batchDocs := range gox.Batch(docs, BatchSize) {
			if _, err := dst.InsertMany(ctx, batchDocs); err != nil {
				return fmt.Errorf("insert batch failed: %w", err)
			}
		}

		return nil
	}

	wrap := func(ctx mongo.SessionContext) (any, error) {
		return nil, callback(ctx)
	}

	if _, err := sess.WithTransaction(ctx, wrap); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}

func BackupMongoData(
	ctx context.Context,
	col *mongo.Collection,
	filter bson.D,
	backupColName string,
) (err error) {
	backupCol := col.Database().Collection(backupColName)

	count, err := backupCol.CountDocuments(ctx, bson.D{})
	if err != nil {
		return fmt.Errorf("failed to check existing backup: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("backup collection %s already has data", backupColName)
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to fetch documents for backup: %w", err)
	}

	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var docs []any

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return fmt.Errorf("failed to decode document: %w", err)
		}

		docs = append(docs, doc)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error during backup: %w", err)
	}

	if len(docs) == 0 {
		return nil
	}

	if _, err := backupCol.InsertMany(ctx, docs); err != nil {
		return fmt.Errorf("failed to insert backup documents: %w", err)
	}

	return nil
}
