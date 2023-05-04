package config

import (
	"context"
	"fmt"
	"time"

	"github.com/ilhm-rai/mygram/entity"
	"github.com/ilhm-rai/mygram/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func NewPostgresDatabase(configuration Config) *gorm.DB {
	ctx, cancel := NewPostgresContext()
	defer cancel()

	host := configuration.Get("PGHOST")
	user := configuration.Get("PGUSER")
	password := configuration.Get("PGPASSWORD")
	dbname := configuration.Get("PGDATABASE")
	port := configuration.Get("PGPORT")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{TranslateError: true})

	exception.PanicIfNeeded(err)

	DB.Debug().AutoMigrate(entity.User{}, entity.Comment{}, entity.Photo{}, entity.SocialMedia{})

	DB.WithContext(ctx)

	return DB
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetDB() *gorm.DB {
	return DB
}
