package server

import (
	"github.com/Imtiaz246/Thesis-Management-System/docs"
	apiV1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/handler"
	"github.com/Imtiaz246/Thesis-Management-System/internal/middleware"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/server/http"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *token.JWT,
	userHandler *handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Hello world!",
		})
	})

	v1 := s.Group("/api/v1")
	v1.POST("/login", userHandler.Login)
	{
		student := v1.Group("students")
		student.POST("/request-register", userHandler.ReqRegister)
		student.POST("/verify-email", userHandler.VerifyEmail)
		student.POST("/register", userHandler.Register)
	}

	return s
}
