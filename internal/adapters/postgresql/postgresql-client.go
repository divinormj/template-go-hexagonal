package postgresql

import (
	"time"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
)

func NewPostgressqlClient() {
	logger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)
	db, err := gorm.Open(postgres.Open("dsn"), &gorm.Config{Logger: logger})

	if err != nil {
		panic("failed to connect database")
	}

	// Adiciona o plugin opentelemetry ao GORM
	err = db.Use(opentelemetry.New(opentelemetry.Config{
		TracerProvider: trace.TracerProvider(), // Ou seu próprio TracerProvider
	}))
	if err != nil {
		panic("failed to add opentelemetry plugin")
	}
}
