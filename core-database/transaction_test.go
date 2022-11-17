package core_database_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	core_database "github.com/SimifiniiCTO/core/core-database"
)

var (
	dbConn       *core_database.DatabaseConn
	databaseType = "postgres"
	host         = "localhost"
	port         = 5433
	user         = "postgres"
	password     = "postgres"
	dbname       = "postgres"
	globalCtx    = context.TODO()
)

func init() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	dbConn = core_database.NewDatabaseConn(connectionString, databaseType, 500*time.Millisecond)
}

// TestTransaction Tests the result of a transaction
func TestTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_TransactionPassed", TransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_Transaction", TransactionFailed)
}

func TestComplexTransaction(t *testing.T) {
	// success scenarios
	t.Run("TestName:Passed_ComplexTransaction", ComplexTransactionPassed)
	// failure scenarios
	t.Run("TestName:Failed_ComplexTransaction", ComplexTransactionFailed)
}

func TransactionFailed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) error {
		return errors.New("failed transactions")
	}

	err := dbConn.PerformTransaction(globalCtx, f)
	assert.NotEmpty(t, err)
}

func ComplexTransactionFailed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		return nil, errors.New("failed transactions")
	}

	output, err := dbConn.PerformComplexTransaction(globalCtx, f)
	assert.NotEmpty(t, err)
	assert.Empty(t, output)
}

func TransactionPassed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) error {
		return nil
	}

	err := dbConn.PerformTransaction(globalCtx, f)
	assert.Empty(t, err)
}

func ComplexTransactionPassed(t *testing.T) {
	f := func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		return true, nil
	}

	output, err := dbConn.PerformComplexTransaction(globalCtx, f)
	assert.Empty(t, err)
	assert.NotEmpty(t, output)
}
