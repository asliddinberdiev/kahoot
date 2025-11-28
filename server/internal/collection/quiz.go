package collection

import (
	"context"
	"log"

	"github.com/asliddinberdiev/kahoot/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuizCollection struct {
	collection *mongo.Collection
}

func Quiz(collection *mongo.Collection) *QuizCollection {
	return &QuizCollection{
		collection: collection,
	}
}

func (c QuizCollection) InsertQuiz(ctx context.Context, quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(ctx, quiz)
	return err
}

func (c QuizCollection) GetQuizById(ctx context.Context, id primitive.ObjectID) (*entity.Quiz, error) {
	result := c.collection.FindOne(ctx, bson.M{"_id": id})

	var quiz entity.Quiz
	if err := result.Decode(&quiz); err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (c QuizCollection) GetQuizzes(ctx context.Context) ([]entity.Quiz, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	quizzes := make([]entity.Quiz, 0)
	if err = cursor.All(ctx, &quizzes); err != nil {
		log.Println(err)
		return nil, err
	}

	return quizzes, nil
}
