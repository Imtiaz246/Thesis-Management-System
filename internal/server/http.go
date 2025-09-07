package server

import (
	"github.com/Imtiaz246/Thesis-Management-System/docs"
	"github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
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
	batchHandler *handler.BatchHandler,
	teamHandler *handler.TeamHandler,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(logger, conf),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		v1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Hello world!",
		})
	})

	apiv1 := s.Group("/api/v1")
	apiv1.POST("/login", userHandler.Login)
	{
		students := apiv1.Group("students")
		students.POST("/request-register", userHandler.ReqRegister)
		students.POST("/verify-email", userHandler.VerifyEmail)
		students.POST("/register", userHandler.Register)
	}
	{
		users := apiv1.Group("users")
		users.GET("/:uni_id/profile", middleware.NoStrictAuth(jwt, logger), userHandler.GetProfile)
		users.PUT("/profile", middleware.StrictAuth(jwt, logger), userHandler.UpdateProfile)
	}
	{
		batchGroup := apiv1.Group("batches")
		batchGroup.GET("/", batchHandler.ListBatch)
		batchGroup.POST("/", middleware.StrictAuth(jwt, logger), batchHandler.CreateBatch)
		batchGroup.GET("/open", batchHandler.ListOpenBatches)
		batchGroup.GET("/:id", batchHandler.GetBatch)
		batchGroup.PUT("/:id", middleware.StrictAuth(jwt, logger), batchHandler.UpdateBatch)
		batchGroup.DELETE("/:id", middleware.StrictAuth(jwt, logger), batchHandler.DeleteBatch)
		batchGroup.POST("/:id/register", middleware.StrictAuth(jwt, logger), batchHandler.RegisterToBatch)
		batchGroup.GET("/:id/registers", middleware.StrictAuth(jwt, logger), batchHandler.ListBatchRegisters)
		batchGroup.PUT("/:id/close", middleware.StrictAuth(jwt, logger), batchHandler.CloseBatch)

		// TODO: separate team related routes from batch routes
		{
			teamGroup := batchGroup.Group("/:id/teams", middleware.StrictAuth(jwt, logger))
			teamGroup.POST("/", teamHandler.CreateTeam)                                                        // student has to be registered in batch
			teamGroup.GET("/joined", teamHandler.GetJoinedTeam)                                                // student has to be team member
			teamGroup.PUT("/:team_id/leave", teamHandler.LeaveTeam)                                            // team member
			teamGroup.POST("/:team_id/invitations/send_to/:target_student_uni_id", teamHandler.SendInvitation) // requester has to be team member, target user has to be registered in batch
			teamGroup.PUT("/:team_id/invitations/:invitation_id/reject", teamHandler.RejectInvitation)         // team invitation has to exits or no exits always pass
			teamGroup.PUT("/:team_id/invitations/:invitation_id/accept", teamHandler.AcceptInvitation)         // team invitation has to exists
			teamGroup.GET("/:team_id/invitations", teamHandler.ListInvitations)                                // all team invitation

			// TODO: teacher selection specific routers
		}
	}

	return s
}
