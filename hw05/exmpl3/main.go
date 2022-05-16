package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" // вместо github.com/gorilla/mux
)

func main() {
	// объявляем mux.Router
	router := chi.NewRouter()
	// регистрируем анонимную функцию-обработчик на корневой маршрут для метода GET
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET HANDLER")
	})
	// и для метода POST
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST HANDLER")
	})
	// вызов mux.Vars["id"] вернет параметр, переданный URI запроса на месте {id}

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Fprintln(w, "GET BY ID HANDLER. RESOURCE ID IS", id)
		return
	})
	// в данном случае, помимо параметра id, mux.Vars["name"] вернет параметр, переданный URI
	//запроса на месте {name}

	router.Get("/{id}/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		id, name := chi.URLParam(r, "id"), chi.URLParam(r, "name")
		fmt.Fprintf(w, "GET BY ID HANDLER WITH NAME. RESOURCE ID IS %s AND NAME IS %s\n", id,
			name)
		return
	})
}
