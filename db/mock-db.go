package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/gomega"
	slogGorm "github.com/orandin/slog-gorm"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSQLMockGorm() (sqlMock sqlmock.Sqlmock, gormDb *gorm.DB) {
	var (
		db  *sql.DB
		err error
	)

	db, sqlMock, err = sqlmock.New()
	Ω(err).Should(Succeed())

	gormDb, err = gorm.Open(
		pg.New(pg.Config{Conn: db}),
		&gorm.Config{Logger: slogGorm.New()},
	)
	Ω(err).Should(Succeed())

	return sqlMock, gormDb
}
