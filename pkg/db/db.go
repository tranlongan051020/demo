package db

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Location string
}

func connectDB(config *DatabaseConfig) (gormDB *gorm.DB) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		config.Host,
		config.Port,
		config.DBName,
		config.User,
		config.Password,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sql, err := gormDB.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	maximumIdleConnection := 10
	sql.SetMaxIdleConns(maximumIdleConnection)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maximumOpenConnection := 20
	sql.SetMaxOpenConns(maximumOpenConnection)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sql.SetConnMaxLifetime(time.Hour)

	return gormDB
}
