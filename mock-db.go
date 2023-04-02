package test_utils

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetMockGorm() (sqlMock sqlmock.Sqlmock, gormDb *gorm.DB, err error) {
	var db *sql.DB

	db, sqlMock, err = sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDb, err = gorm.Open(
		pg.New(pg.Config{Conn: db}), &gorm.Config{
			Logger: logger.New(
				&TestLogger,
				logger.Config{
					SlowThreshold:             200 * time.Millisecond, //nolint:gomnd
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: false,
					Colorful:                  true,
				},
			),
		},
	)

	return sqlMock, gormDb, err
}
