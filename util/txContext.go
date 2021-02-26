package util

import "database/sql"

type TxContext struct {
	TX, _ *sql.Tx
}
