package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"
	slogGorm "github.com/orandin/slog-gorm"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockForGorm(dialector gorm.Dialector) *gorm.DB {
	var gormDb, err = gorm.Open(dialector, &gorm.Config{Logger: slogGorm.New()})

	gomega.Expect(err).To(gomega.Succeed())

	return gormDb
}

func NewMockForGormPostgreSQL() (sqlmock.Sqlmock, *gorm.DB) {
	var db, sqlMock = NewSQLMock()

	return sqlMock, NewMockForGorm(pg.New(pg.Config{Conn: db}))
}
