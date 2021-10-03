package ramsqldb

import "testing"

func TestOpenTestDb(t *testing.T) {
	db := OpenTestDb(t)
	if db == nil {
		t.Fatal("OpenTestDb() returned nil")
	}
}
