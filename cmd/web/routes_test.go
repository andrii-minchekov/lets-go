package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/andrii-minchekov/lets-go/app/impl"
	"github.com/andrii-minchekov/lets-go/app/impl/cfg"
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/andrii-minchekov/lets-go/mocks"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func Test_CreateSnippetSuccessUT(t *testing.T) {
	inputSnippet := snippet()
	snippetJson, _ := json.Marshal(inputSnippet)
	r := httptest.NewRequest("Post", "http://example.com", bytes.NewReader(snippetJson))
	w := httptest.NewRecorder()
	var app = &App{Cases: mockCreateSnippet(inputSnippet)}

	app.CreateSnippetJson(w, r)

	require.Equal(t, `{"id": 1}`, string(w.Body.Bytes()))
	require.Equal(t, 201, w.Result().StatusCode)
}

func snippet() snp.Snippet {
	rand.Seed(time.Now().UnixNano())
	inputSnippet := snp.Snippet{Title: strconv.Itoa(rand.Int()), Content: fmt.Sprintf("Content %d", rand.Int())}
	return inputSnippet
}

func mockCreateSnippet(snippet snp.Snippet) *mocks.UseCases {
	cases := mocks.UseCases{}
	cases.On("CreateSnippet", snippet).Return(int64(1), nil)
	return &cases
}

func Test_SignUpSuccessIT(t *testing.T) {
	srv := httptest.NewServer((&App{Config: cfg.FlagConfig, Cases: impl.NewComposedUseCases(cfg.FlagConfig)}).Routes())
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
