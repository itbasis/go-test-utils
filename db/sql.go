package db

import (
	"database/sql"

	goSqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"
)

func NewSQLMock() (*sql.DB, goSqlMock.Sqlmock) {
	var db, sqlMock, err = goSqlMock.New()

	gomega.Expect(err).To(gomega.Succeed())

	return db, sqlMock
}
