package service

import (
	"context"

	"github.com/asliddinberdiev/kahoot/internal/collection"
	"github.com/asliddinberdiev/kahoot/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) GetQuizzes(ctx context.Context) ([]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes(ctx)
}
