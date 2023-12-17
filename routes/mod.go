package routes

import (
    "net/http"
    "github.com/gorilla/sessions"
)

// TODO
var Sessions = sessions.NewCookieStore([]byte("abc123"))

// GET /
func Home(w http.ResponseWriter, r *http.Request) {
}

// GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
}

// POST /signup
func SignupPost(w http.ResponseWriter, r *http.Request) {
}

// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
}

// POST /login
func LoginPost(w http.ResponseWriter, r *http.Request) {
}
