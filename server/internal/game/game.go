package game

import (
	"math/rand"
	"strconv"

	"github.com/asliddinberdiev/kahoot/internal/entity"
	"github.com/google/uuid"
)

type Game struct {
	Id   uuid.UUID
	Quiz entity.Quiz
	Code string
}

func generateCode() string {
	return strconv.Itoa(100_000 + rand.Intn(900_000))
}

func New(quiz entity.Quiz) Game {
	return Game{
		Id:   uuid.New(),
		Quiz: quiz,
		Code: generateCode(),
	}
}

func (g *Game) Start() {

}
