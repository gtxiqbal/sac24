package helper

import "database/sql"

func RowsClose(rows *sql.Rows) {
	err := rows.Close()
	PanicIfError(err)
}

func StmtClose(stmt *sql.Stmt) {
	err := stmt.Close()
	PanicIfError(err)
}
