package core_database

import (
	"errors"
	"os"
	"time"

	"github.com/giantswarm/retry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConn struct {
	Engine *gorm.DB
}

// NewDatabaseConn obtains a reference to a database connection object
func NewDatabaseConn(connString, databaseType string) *DatabaseConn {
	if connString == "" {
		// crash the process
		os.Exit(1)
	}

	conn, err := Connect(connString, databaseType)
	if err != nil {
		panic("failed to connect to database")
	}

	return &DatabaseConn{Engine: conn}
}

// Connect attempts to connect to the database using retries
func Connect(connectionString, databaseType string) (*gorm.DB, error) {
	var connection = make(chan *gorm.DB, 1)

	err := retry.Do(
		func(conn chan<- *gorm.DB) func() error {
			return func() error {
				newConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
				conn <- newConn
				return err
			}
		}(connection),
		retry.MaxTries(5),
		retry.Timeout(time.Second*10),
		retry.Sleep(1*time.Second),
	)

	if err != nil {
		return nil, errors.New("exceeded retries")
	}

	return <-connection, nil
}
