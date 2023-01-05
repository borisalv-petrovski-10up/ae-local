package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/spf13/viper"

	"source.cloud.google.com/sk-borislav/ae-local/handler"
	"source.cloud.google.com/sk-borislav/ae-local/repository"
)

type AppConfig struct {
	Port                  string `mapstructure:"PORT"`
	ProjectID             string `mapstructure:"PROJECT_ID"`
	DatastoreEmulatorHost string `mapstructure:"DATASTORE_EMULATOR_HOST"`
}

func main() {
	ctx := context.Background()

	cfg := AppConfig{
		Port:                  os.Getenv("PORT"),
		ProjectID:             os.Getenv("PROJECT_ID"),
		DatastoreEmulatorHost: os.Getenv("DATASTORE_EMULATOR_HOST"),
	}

	// in case running on local machine without app engine
	if cfg.Port == "" || cfg.ProjectID == "" || cfg.DatastoreEmulatorHost == "" {
		err := loadConfig(&cfg)
		if err != nil {
			panic(err)
		}
		_ = os.Setenv("DATASTORE_EMULATOR_HOST", cfg.DatastoreEmulatorHost)
	}

	datastoreClient, err := datastore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(datastoreClient)
	h := handler.NewHandler(repo)

	http.HandleFunc("/create-article", h.CreateArticle)
	http.HandleFunc("/get-article/", h.GetArticle)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func loadConfig(cfg *AppConfig) error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&cfg)
}
