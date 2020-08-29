package model

import "ark/db"

func SyncDB() error {
	return db.DBE.Sync2( new(User), new(Role))
}
