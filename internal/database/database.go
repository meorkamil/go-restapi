package database

import (
	"fmt"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"io"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	lg = util.NewLogger()
	db *gorm.DB
)

type dbConfig struct {
	Host    string
	User    string
	Pass    string
	Port    string
	DBName  string
	DBFlags string
	Type    string
}

func InitDB(c *model.Config) *dbConfig {
	return &dbConfig{
		Host:    c.Database.Host,
		User:    c.Database.User,
		Pass:    c.Database.Pass,
		DBName:  c.Database.DBName,
		DBFlags: c.Database.DBFlags,
		Type:    c.Database.Type,
	}
}

// Create connection
func (c *dbConfig) Con() (*gorm.DB, error) {
	// Set default logger
	sqlLog := logger.New(
		log.New(io.Discard, "", 0),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
		},
	)

	// Create connection based on database type
	switch {
	case c.Type == "mysql":
		/// Do mysql connection proceed with postgres for now
		db, err := gorm.Open(
			mysql.Open(c.dsn()),
			&gorm.Config{
				Logger: sqlLog,
			},
		)
		if err != nil {
			lg.Fatal("Database", err)
			return nil, fmt.Errorf("Database %s", err)
		}
		return db, nil
	default:
		db, err := gorm.Open(
			postgres.Open(c.dsn()),
			&gorm.Config{
				Logger: sqlLog,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("Database %s", err)
		}
		return db, nil
	}
}

// Create automication GORM
func (c *dbConfig) Migration() error {
	// Run AutoMigrate
	m, err := c.Con()
	if err != nil {
		return fmt.Errorf("Database %s", err)
	}

	m.AutoMigrate(
		&model.Employee{},
		&model.Product{},
	)

	lg.Info("GORM AutoMigrate Completed")

	return nil
}

// Create connection string
func (c *dbConfig) dsn() string {
	// Return a string of dsn for open a connection
	switch {
	case c.Type == "mysql":
		// Create mysql connection string
		return c.User + ":" + c.Pass + "@tcp(" + c.Host + c.Port + ")/" + c.DBName + c.DBFlags
	default:
		return "host=" + c.Host + " password=" + c.Pass + " user=" + c.User + " dbname=" + c.DBName + " " + c.DBFlags
	}
}
