package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	errRecover := recover()
	if errRecover == nil {
		err := tx.Commit()
		PanicIfError(err)
	} else {
		err := tx.Rollback()
		PanicIfError(err)
		panic(errRecover)
	}
}
