package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/newUser1337/task-news/internal/config"
	"github.com/newUser1337/task-news/internal/context"
	"github.com/newUser1337/task-news/internal/migration"
	"github.com/newUser1337/task-news/internal/repository"
	"github.com/newUser1337/task-news/internal/router"
	"github.com/newUser1337/task-news/internal/server"
	"github.com/newUser1337/task-news/internal/service/news"
	"github.com/newUser1337/task-news/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildApp() {
	// Config
	cfg := config.GetConfig()
	// Context
	appCtx, err := context.NewAppCtx()
	if err != nil {
		log.Fatalf("failed to create context %s", err)
	}
	// Creating mongo connection
	mongoClient, err := mongo.Connect(appCtx.GetGoContext(), options.Client().ApplyURI(cfg.Core.Mongo.Address))
	if err != nil {
		log.Fatalf("failed to connect to mongo db: %s", err)
	}
	// migration
	if err := migration.Migrate(cfg.Core.Mongo); err != nil {
		log.Fatalf("failed to migrate mongo %s", err)
	}
	// Repository access to db data
	rep := repository.NewRepository(mongoClient)
	// Usecase - main logic
	usc := usecase.NewUsecase(rep)
	// Router register endpoint and handlers
	r := router.NewRouter(appCtx.GetGoContext(), usc)
	// Build server
	srv := server.NewServer(cfg.Core.Port, r)

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	news.StartNewNewsServer(appCtx, cfg, rep)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	<-signalCh
	if err := srv.Stop(); err != nil {
		log.Printf("incorrect shutdown http server occured %s\n", err)
	}
	appCtx.Cancel()
	appCtx.WgWait()

	if err := mongoClient.Disconnect(appCtx.GetGoContext()); err != nil {
		log.Print("failed to disconnect mongo db")
	}

	log.Println("shutting down")
	os.Exit(0)
}
