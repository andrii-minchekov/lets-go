package main

import (
	"flag"
	"github.com/alexedwards/scs"
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlerFuncIntg(t *testing.T) {
	dsn := flag.String("dsn", "user=postgres dbname=snippetbox password=postgres sslmode=disable", "Postgres Datasource Name")
	secret := flag.String("secret", "s6Nd%+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

	flag.Parse()

	db := connect(*dsn)
	sessionManager := scs.NewCookieManager(*secret)
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)

	//defer db.Close()
	app := &App{
		Database: &models.Database{db},
		Sessions: sessionManager,
	}

	if testing.Short() {
		println("skipping")
		t.Skip()
	}

	reqStr := `{"Title":"e2eTitle", "Content": "e2eContent", "Expires": "17.01.2021"}`
	req, err := http.NewRequest("POST", "/snippets/new", strings.NewReader(reqStr))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CreateSnippet)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
