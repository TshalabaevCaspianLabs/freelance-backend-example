package main

import (
    "log"
    "myapp/internal/router"
    "net/http"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    _ "myapp/docs" 
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @title Freelance Backend API
// @version 1.0
// @host localhost:8080/api/v1
// @BasePath /

func main() {
    dsn := "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    r := router.SetupRouter(db)

    // Добавление маршрута для Swagger документации
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
