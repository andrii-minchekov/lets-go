package web

import (
	"encoding/json"
	"fmt"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"io/ioutil"
	"net/http"
)

func (app *App) SignUpUserJson(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.ServerError(w, err)
	}
	var user = &usr.User{}
	err = json.Unmarshal(bytes, user)

	id, err := app.Cases.SignupUser(*user)

	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, `{"id": %d}`, id)
	if err != nil {
		app.ServerError(w, err)
	}
}
