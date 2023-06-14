package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	itbasisTestUtils "github.com/itbasis/go-test-utils/v2"
	. "github.com/onsi/gomega"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func GetSqlMockGorm() (sqlMock sqlmock.Sqlmock, gormDb *gorm.DB) {
	var (
		db  *sql.DB
		err error
	)

	db, sqlMock, err = sqlmock.New()
	Ω(err).Should(Succeed())

	logger := zapgorm2.New(itbasisTestUtils.TestLogger)
	logger.SetAsDefault()

	gormDb, err = gorm.Open(
		pg.New(pg.Config{Conn: db}),
		&gorm.Config{Logger: logger.LogMode(gormLogger.Info)},
	)
	Ω(err).Should(Succeed())

	return sqlMock, gormDb
}
