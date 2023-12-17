package main

import (
    //"fmt"
    "net/http"

    "artbooru/routes"
    "artbooru/db"

    "github.com/julienschmidt/httprouter"
)

func main() {
    db_file := "artbooru.sqlite"
    db.Connect(db_file)

    router := httprouter.New()
    router.HandlerFunc("GET", "/", routes.Home)
    router.HandlerFunc("GET", "/signup", routes.Home)
    router.HandlerFunc("POST", "/signup", routes.Home)
    router.HandlerFunc("GET", "/login", routes.Home)
    router.HandlerFunc("POST", "/login", routes.Home)

    fs := http.FileServer(http.Dir("images/"))
    router.Handler("GET", "/images/", http.StripPrefix("/images/", fs))

    if err := http.ListenAndServe(":3000", router); err != nil {
        panic(err)
    }
}
