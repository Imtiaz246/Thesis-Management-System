package server

import (
	"context"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}

func (m *Migrate) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		m.log.Error("Failed to migrate User", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Student{}); err != nil {
		m.log.Error("Failed to migrate Student", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Teacher{}); err != nil {
		m.log.Error("Failed to migrate Teacher", zap.Error(err))
		return err
	}
	if err := m.db.AutoMigrate(&model.Stuff{}); err != nil {
		m.log.Error("Failed to migrate Stuff")
	}
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}

func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}
