package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/CallumKerrEdwards/library-api/pkg/books"
)

func (d *Database) GetBook(ctx context.Context, bookID string) (books.Book, error) {
	booksCollection := d.Client.Database("library").Collection("books")
	filter := bson.D{{Key: "id", Value: bookID}}

	var result books.Book
	err := booksCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		d.Logger.WithError(err).Errorln("Cannot get book wit id", bookID)
		return books.Book{}, err
	}

	return result, nil
}

func (d *Database) GetAllBooks(ctx context.Context) ([]books.Book, error) {
	booksCollection := d.Client.Database("library").Collection("books")

	var results []books.Book

	cursor, err := booksCollection.Find(ctx, bson.D{{}})
	if err != nil {
		d.Logger.WithError(err).Errorln("Cannot get all items from collection")
		return []books.Book{}, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem books.Book

		err := cursor.Decode(&elem)
		if err != nil {
			d.Logger.WithError(err).Errorln("Cannot decode element into book")
			return []books.Book{}, err
		} else {
			results = append(results, elem)
		}
	}

	if err := cursor.Err(); err != nil {
		d.Logger.WithError(err).Errorln("Error reading cursor")
		return []books.Book{}, err
	}

	return results, nil
}

func (d *Database) PostBook(ctx context.Context, bookToInsert *books.Book) (books.Book, error) {
	booksCollection := d.Client.Database("library").Collection("books")
	result, err := booksCollection.InsertOne(ctx, bookToInsert)

	if err != nil {
		d.Logger.WithError(err).Errorln("Cannot store book with id", bookToInsert.ID)
		return books.Book{}, err
	}

	d.Logger.WithField("mongodb_id", result.InsertedID).Debugln("Successfully stored book with id", bookToInsert.ID)

	return *bookToInsert, nil
}

func (d *Database) UpdateBook(ctx context.Context, bookID string, updatedBook *books.Book) (bool, books.Book, error) {
	booksCollection := d.Client.Database("library").Collection("books")

	filter := bson.D{{Key: "id", Value: bookID}}

	updateResult, err := booksCollection.ReplaceOne(ctx, filter, updatedBook)
	if err != nil {
		d.Logger.WithError(err).Errorln("Cannot update book with id", bookID)
		return false, books.Book{}, err
	}

	d.Logger.WithField("mongodb_id", updateResult.UpsertedID).Debugln("Successfully updated book with id", bookID)

	return updateResult.MatchedCount == 1, *updatedBook, nil
}

func (d *Database) DeleteBook(ctx context.Context, bookID string) error {
	booksCollection := d.Client.Database("library").Collection("books")

	filter := bson.D{{Key: "id", Value: bookID}}

	_, err := booksCollection.DeleteOne(ctx, filter)

	return err
}
