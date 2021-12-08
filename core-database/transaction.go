package core_database

import (
	"context"

	"gorm.io/gorm"
)

// Tx is a type serving as a function decorator for common database transactions
type Tx func(ctx context.Context, tx *gorm.DB) error

// CmplxTx is a type serving as a function decorator for complex database transactions
type CmplxTx func(ctx context.Context, tx *gorm.DB) (interface{}, error)

// PerformTransaction takes as input an anonymous function witholding
// logic to perform within a transaction. This function is then invoked within a transaction.
// if unsuccessful or any error is raised throughout the transaction, then, the transaction
// is rolled back. Returned is any error occuring throughout the transaction lifecycle
func (db *DatabaseConn) PerformTransaction(ctx context.Context, transaction Tx) error {
	f := func(tx *gorm.DB) error {
		return transaction(ctx, tx)
	}

	return db.Engine.Transaction(f)
}

// PerformComplexTransaction takes as input an anonymous function witholding logic
// to perform within a transaction returning an abstract type. This function is then invoked within a transaction
// and depending on the occurrence of any specific errors, the transaction is either committed to the database
// or completely rolled back. This returns the result obtained from the invocation of the anonymous function as
// well as any error occuring throughout the transaction lifecycle.
func (db *DatabaseConn) PerformComplexTransaction(ctx context.Context, transaction CmplxTx) (interface{}, error) {
	tx := db.Engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	result, err := transaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, tx.Commit().Error
}
