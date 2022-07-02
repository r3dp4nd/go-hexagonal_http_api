package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/platform/server"
	"github.com/Sraik25/go-hexagonal_http_api/01-03-architectured-gin-healthcheck/internal/platform/storage/storagemocks"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := storagemocks.NewCourseRepository(*db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
