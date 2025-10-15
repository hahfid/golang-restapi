package di

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang-restapi/infra/database"
	appconfig "golang-restapi/internal/app/config"
	"golang-restapi/internal/app/feature/auth/user/domain"
	userhandler "golang-restapi/internal/app/feature/auth/user/handler"
	userquery "golang-restapi/internal/app/feature/auth/user/query"
	userrepository "golang-restapi/internal/app/feature/auth/user/repository"
	userroutes "golang-restapi/internal/app/feature/auth/user/routes"
	bookdomain "golang-restapi/internal/app/feature/library/book/domain"
	bookhandler "golang-restapi/internal/app/feature/library/book/handler"
	bookquery "golang-restapi/internal/app/feature/library/book/query"
	bookrepository "golang-restapi/internal/app/feature/library/book/repository"
	bookroutes "golang-restapi/internal/app/feature/library/book/routes"
	"golang-restapi/internal/app/middleware"
	"golang-restapi/pkg/jwt"
	"gorm.io/gorm"
)

type Container struct {
	engine *gin.Engine
	db     *gorm.DB
	cfg    appconfig.Config
}

func NewContainer() (*Container, error) {
	cfg := appconfig.Load()

	db, err := database.NewMySQLConnection(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	if err := db.AutoMigrate(&domain.User{}, &bookdomain.Book{}); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	engine := gin.Default()

	tokenManager := jwt.NewTokenManager(cfg.Security.JWTSecret)
	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	userQuery := userquery.NewUserQuery(db)
	userRepo := userrepository.NewGormRepository(userQuery)
	authService := userhandler.NewAuthService(userRepo, tokenManager)
	authHandler := userhandler.NewAuthHandler(authService)
	userroutes.RegisterRoutes(engine, authHandler)

	bookQuery := bookquery.NewBookQuery(db)
	bookRepo := bookrepository.NewGormRepository(bookQuery)
	bookHandler := bookhandler.NewBookHandler(bookRepo)
	api := engine.Group("/api")
	api.Use(authMiddleware.Handler())
	bookroutes.RegisterRoutes(api, bookHandler)

	return &Container{
		engine: engine,
		db:     db,
		cfg:    cfg,
	}, nil
}

func (c *Container) Run() error {
	return c.engine.Run(":" + c.cfg.Server.Port)
}

func (c *Container) Engine() *gin.Engine {
	return c.engine
}

func (c *Container) DB() *gorm.DB {
	return c.db
}
