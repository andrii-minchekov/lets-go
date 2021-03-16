package main

import (
	"fmt"
	"github.com/andrii-minchekov/lets-go/domain/snippet"
	"github.com/andrii-minchekov/lets-go/domain/user"
	"net/http"
	"strconv"

	"github.com/andrii-minchekov/lets-go/cmd/web/forms"
)

// Home Display a "Hello from Snippetbox" message
func (app *App) Home(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	snippets, err := app.Cases.LatestSnippets()

	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "home.page.html", &HTMLData{
		Snippets: snippets,
	})

}

// ShowSnippet ...
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.URL.Query().Get(":id"), 10, 64)

	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Cases.GetSnippet(id)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	if snippet == nil {
		app.NotFound(w)
		return
	}

	session := app.Sessions.Load(r)
	flash, err := session.PopString(w, "flash")

	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "show.page.html", &HTMLData{
		Snippet: snippet,
		Flash:   flash,
	})
}

// NewSnippetView ...
func (app *App) NewSnippetView(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "new.page.html", &HTMLData{
		Form: &forms.NewSnippet{},
	})
}

// CreateSnippet ...
func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.NewSnippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: r.PostForm.Get("expires"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "new.page.html", &HTMLData{Form: form})
		return
	}

	id, err := app.Cases.CreateSnippet(snp.Snippet{
		Title:   form.Title,
		Content: form.Content,
	})

	if err != nil {
		app.ServerError(w, err)
		return
	}

	session := app.Sessions.Load(r)

	err = session.PutString(w, "flash", "Your snippet was saved!")

	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}

// SignupUserView ...
func (app *App) SignupUserView(w http.ResponseWriter, r *http.Request) {
	app.RenderHTML(w, r, "signup.page.html", &HTMLData{
		Form: &forms.SignupUser{},
	})
}

// SignupUser ...
func (app *App) SignupUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.SignupUser{
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "signup.page.html", &HTMLData{Form: form})
		return
	}

	_, err = app.Cases.SignupUser(usr.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	})

	if err != nil {
		app.ServerError(w, err)
		return
	}

	msg := "Your signup was successful. Please log in using your credentials."
	session := app.Sessions.Load(r)

	err = session.PutString(w, "flash", msg)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

// SigninUserView ...
func (app *App) SigninUserView(w http.ResponseWriter, r *http.Request) {

	session := app.Sessions.Load(r)

	flash, err := session.PopString(w, "flash")
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "login.page.html", &HTMLData{
		Flash: flash,
		Form:  &forms.LoginUser{},
	})
}

// SignInUser ...
func (app *App) SignInUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.ClientError(w, http.StatusBadRequest)
		return
	}

	form := &forms.LoginUser{
		Email:    r.PostForm.Get("email"),
		Password: r.PostForm.Get("password"),
	}

	if !form.Valid() {
		app.RenderHTML(w, r, "login.page.html", &HTMLData{Form: form})
		return
	}

	currentUserID, err := app.Cases.SignInUser(form.Email, form.Password)

	if err == usr.ErrInvalidCredentials {
		form.Failures["Generic"] = "Email or Password is incorrect"
		app.RenderHTML(w, r, "login.page.html", &HTMLData{Form: form})
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}

	// Add the ID of the current user to the session
	session := app.Sessions.Load(r)
	err = session.PutInt(w, "currentUserID", int(currentUserID))

	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, "/snippet/new", http.StatusSeeOther)

}

// LogoutUser ...
func (app *App) LogoutUser(w http.ResponseWriter, r *http.Request) {

	// Remove the currentUserID from the session data.
	session := app.Sessions.Load(r)
	err := session.Remove(w, "currentUserID")

	if err != nil {
		app.ServerError(w, err)
		return

	}

	// Redirect the user to the homepage.
	http.Redirect(w, r, "/", 303)
}
