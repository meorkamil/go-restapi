package database

import (
	"go-restapi/internal/model"
	"go-restapi/internal/util"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
func (c *dbConfig) Con() *gorm.DB {
	// Create connection based on database type
	switch {
	case c.Type == "mysql":
		/// Do mysql connection proceed with postgres for now
		db, err := gorm.Open(
			mysql.Open(c.dsn()),
			&gorm.Config{},
		)
		if err != nil {
			lg.Fatal("GORM open database connection failed:", err)
		}
		return db
	default:
		db, err := gorm.Open(
			postgres.Open(c.dsn()),
			&gorm.Config{},
		)
		if err != nil {
			lg.Fatal("GORM open database connection failed:", err)
		}
		return db
	}
	return db
}

// Create automication GORM
func (c *dbConfig) Migration() error {
	// Run AutoMigrate
	if err := c.Con().AutoMigrate(
		&model.Employee{},
		&model.Product{},
	); err != nil {
		return err
	}

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
