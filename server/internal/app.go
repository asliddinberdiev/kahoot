package internal

import (
	"context"
	"log"
	"time"

	"github.com/asliddinberdiev/kahoot/internal/collection"
	"github.com/asliddinberdiev/kahoot/internal/controller"
	"github.com/asliddinberdiev/kahoot/internal/service"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer *fiber.App
	database   *mongo.Database

	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupDb()
	a.setupServices()
	a.setupHttp()

	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:securepassword123@localhost:27017"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("kahoot")
}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.Ws(a.netService)
	app.Get("/ws", websocket.New(wsController.Ws))

	a.httpServer = app
}
