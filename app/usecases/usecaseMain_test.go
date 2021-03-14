package uc

import (
	"github.com/andrii-minchekov/lets-go/app/impl/db"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	db.CleanupDb()
	os.Exit(t.Run())
}
