package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/newUser1337/task-news/internal/context"
	"github.com/newUser1337/task-news/internal/repository"
	"github.com/newUser1337/task-news/internal/router"
	"github.com/newUser1337/task-news/internal/server"
	"github.com/newUser1337/task-news/internal/service/news"
	"github.com/newUser1337/task-news/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildApp() {
	// Config and context
	appCtx, err := context.NewAppCtx()
	if err != nil {
		log.Fatalf("failed to create context %s", err)
	}
	// Creating mongo connection
	mongoClient, err := mongo.Connect(appCtx.GetGoContext(), options.Client().ApplyURI(appCtx.GetConfig().Mongo.Address))
	if err != nil {
		log.Fatalf("failed to connect to mongo db: %s", err)
	}
	// Repository access to db data
	rep := repository.NewRepository(mongoClient)
	// Usecase - main logic
	usc := usecase.NewUsecase(rep)
	// Router register endpoint and handlers
	r := router.NewRouter(appCtx.GetGoContext(), usc)
	// Build server
	srv := server.NewServer(appCtx.GetConfig().Port, r)

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	news.StartNewNewsServer(appCtx, rep)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	<-signalCh
	appCtx.Cancel()
	appCtx.WgWait()

	if err := mongoClient.Disconnect(appCtx.GetGoContext()); err != nil {
		log.Print("failed to disconnect mongo db")
	}

	log.Println("shutting down")
	os.Exit(0)
}
