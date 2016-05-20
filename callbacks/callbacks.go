// Package callbacks extends default Gorm callbacks such as
// executing callback after create transaction commit.
package callbacks

import "github.com/jinzhu/gorm"

// CreateTransactionCommitCallback defines an interface
// for executing callback before/after commit the create
// transaction.
type CreateTransactionCommitCallback interface {
	BeforeCreateTransactionCommit(scope *gorm.Scope)
	AfterCreateTransactionCommit(scope *gorm.Scope)
}

func beforeCreateTransaction(scope *gorm.Scope) {
	if !scope.HasError() {
		if createTransactionCommitCallback, ok := scope.Value.(CreateTransactionCommitCallback); ok {
			createTransactionCommitCallback.BeforeCreateTransactionCommit(scope)
		}
	}
}

func afterCreateTransaction(scope *gorm.Scope) {
	if !scope.HasError() {
		if createTransactionCommitCallback, ok := scope.Value.(CreateTransactionCommitCallback); ok {
			createTransactionCommitCallback.AfterCreateTransactionCommit(scope)
		}
	}
}

// RegisterCallbacks registers custom Gorm callbacks.
func RegisterCallbacks(db *gorm.DB) {
	callback := db.Callback()

	callback.Create().Before("gorm:commit_or_rollback_transaction").Register("callbacks:before_create_transaction", beforeCreateTransaction)
	callback.Create().After("gorm:commit_or_rollback_transaction").Register("callbacks:after_create_transaction", afterCreateTransaction)
}
