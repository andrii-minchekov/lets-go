package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/andrii-minchekov/lets-go/app/impl"
	"github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_SignUpSuccess(t *testing.T) {
	config := impl.NewFlagConfig()
	srv := httptest.NewServer((&App{Config: config, Cases: impl.NewComposedUseCases(config)}).Routes())
	defer srv.Close()

	rand.Seed(time.Now().UTC().UnixNano())
	userJson, _ := json.Marshal(usr.User{Name: "Andrew", Email: fmt.Sprintf("andrew%d@gmail.com", rand.Int()), Password: "12"})
	res, err := http.Post(fmt.Sprintf("%s/users", srv.URL), "application/json", bytes.NewReader(userJson))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("status is %d but want %d", res.StatusCode, 201)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Id int
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Couldn't parse body \"%s\" because \"%s\"", string(body), err.Error())
	}
	require.Greater(t, result.Id, 0)
}
