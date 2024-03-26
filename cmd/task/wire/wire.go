//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Imtiaz246/Thesis-Management-System/internal/server"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/app"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var taskSet = wire.NewSet(server.NewTask)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		taskSet,
		newApp,
	))
}
