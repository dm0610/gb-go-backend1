package main

import (
	"fmt"
	"log"
	"net/http"

	// импорт пакета gorilla/mux
	"github.com/gorilla/mux"
)

func main() {
	// объявляем mux.Router
	router := mux.NewRouter()
	// регистрируем анонимную функцию-обработчик на корневой маршрут для метода GET
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET HANDLER")
	}).Methods(http.MethodGet)

	// вызов mux.Vars["id"] вернет параметр, переданный URI запроса на месте {id}
	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(w, "GET BY ID HANDLER. RESOURCE ID IS", vars["id"])
		return
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "GET BY ID HANDLER WITH NAME. RESOURCE ID IS %s AND NAME IS %s\n",
			vars["id"], vars["name"])
		return
	}).Methods(http.MethodGet)

	// и для метода POST
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST HANDLER")
	}).Methods(http.MethodPost)
	// запускаем сервер, передав в качестве маршрутизатора объект mux.Router
	log.Fatal(http.ListenAndServe(":8020", router))
}
