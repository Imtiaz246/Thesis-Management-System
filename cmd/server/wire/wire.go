//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Imtiaz246/Thesis-Management-System/internal/handler"
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/internal/server"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service"
	batchservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/batch"
	userservice "github.com/Imtiaz246/Thesis-Management-System/internal/service/user"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/app"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/helper/sid"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/mailer"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/server/http"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/token"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewBatchRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	userservice.NewUserService,
	batchservice.NewBatchService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewBatchHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		token.NewJwt,
		mailer.NewMailer,
		newApp,
	))
}
