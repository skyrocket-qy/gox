package envmigrate

// func CopyMongoData(
// 	ctx context.Context,
// 	src, dst *mongo.Collection,
// 	filter bson.D,
// 	replace bool,
// ) error {
// 	cursor, err := src.Find(ctx, filter)
// 	if err != nil {
// 		return fmt.Errorf("failed to find documents in source: %w", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var docs []any
// 	for cursor.Next(ctx) {
// 		var doc bson.M
// 		if err := cursor.Decode(&doc); err != nil {
// 			return fmt.Errorf("failed to decode document: %w", err)
// 		}
// 		delete(doc, "_id") // remove _id to avoid duplicate key error
// 		docs = append(docs, doc)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return fmt.Errorf("cursor error: %w", err)
// 	}

// 	if len(docs) == 0 {
// 		return fmt.Errorf("no documents found in source")
// 	}

// 	if replace {
// 		if _, err := dst.DeleteMany(ctx, filter); err != nil {
// 			return fmt.Errorf("failed to delete existing documents: %w", err)
// 		}
// 	} else {
// 		count, err := dst.CountDocuments(ctx, filter)
// 		if err != nil {
// 			return fmt.Errorf("failed to count destination documents: %w", err)
// 		}
// 		if count > 0 {
// 			return fmt.Errorf("destination collection already contains matching documents")
// 		}
// 	}

// 	sess, err := dst.Database().Client().StartSession()
// 	if err != nil {
// 		return fmt.Errorf("failed to start session: %w", err)
// 	}
// 	defer sess.EndSession(ctx)

// 	callback := func(sc mongo.SessionContext) (any, error) {
// 		const batchSize = 500

// 		for i := 0; i < len(docs); i += batchSize {
// 			end := min(i+batchSize, len(docs))

// 			// insert current batch inside transaction context
// 			if _, err := dst.InsertMany(sc, docs[i:end]); err != nil {
// 				return nil, fmt.Errorf("insert batch failed: %w", err)
// 			}
// 		}
// 		return nil, nil
// 	}

// 	if _, err := sess.WithTransaction(ctx, callback); err != nil {
// 		return fmt.Errorf("transaction error: %w", err)
// 	}

// 	return nil
// }

// func BackupMongoData(
// 	ctx context.Context,
// 	col *mongo.Collection,
// 	filter bson.D,
// 	backupName string,
// ) error {
// 	cursor, err := col.Find(ctx, filter)
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch documents for backup: %w", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var docs []any
// 	for cursor.Next(ctx) {
// 		var doc bson.M
// 		if err := cursor.Decode(&doc); err != nil {
// 			return fmt.Errorf("failed to decode document: %w", err)
// 		}
// 		docs = append(docs, doc)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return fmt.Errorf("cursor error during backup: %w", err)
// 	}

// 	if len(docs) == 0 {
// 		return nil
// 	}

// 	backupCol := col.Database().Collection(backupName)
// 	if _, err := backupCol.InsertMany(ctx, docs); err != nil {
// 		return fmt.Errorf("failed to insert backup documents: %w", err)
// 	}

// 	return nil
// }
