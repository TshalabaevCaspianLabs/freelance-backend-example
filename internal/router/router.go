package router

import (
    "github.com/gin-gonic/gin"
    "myapp/internal/app/handler"
    "myapp/internal/middleware"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // Middleware
    r.Use(middleware.CORSMiddleware())

    v1 := r.Group("/api/v1")
    {
        handler.RegisterUserRoutes(v1, db)
        handler.RegisterNoteRoutes(v1, db)
        handler.RegisterTournamentRoutes(v1, db)
        handler.RegisterMatchRoutes(v1, db)
        handler.RegisterTeamRoutes(v1, db)
        handler.RegisterRatingRoutes(v1, db)
        handler.RegisterBanListRoutes(v1, db)
        handler.RegisterNotificationRoutes(v1, db)
    }

    return r
}
