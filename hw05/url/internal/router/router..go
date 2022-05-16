package router

import (
	"fmt"
	"net/http"
	_ "url/internal/models"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	// объявляем mux.Router
	router := mux.NewRouter()
	// вывод списка из Items с фильтрацией
	router.HandleFunc("/items", ListItemsHandler()).Methods(http.MethodGet)
	// добавление нового Item
	router.HandleFunc("/items", CreateItemHandler()).Methods(http.MethodPost)
	// получение Item по ID
	router.HandleFunc("/items/{id}", GetItemHandler()).Methods(http.MethodGet)
	// изменение Item по ID
	router.HandleFunc("/items/{id}", UpdateItemHandler()).Methods(http.MethodPut)
	// удаление Item по ID
	router.HandleFunc("/items/{id}", DeleteItemHandler()).Methods(http.MethodDelete)
	// загрузка файла изображения
	router.HandleFunc("/items/upload_image", UploadItemImageHandler()).Methods(http.MethodPost)
	// авторизация
	router.HandleFunc("/user/login", LoginHandler()).Methods(http.MethodPost)
	// завершение сессии
	router.HandleFunc("/user/logout", LogoutHandler()).Methods(http.MethodPost)
	return router
}

func ListItemsHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "LIST ALL ITEMS"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func CreateItemHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "CREATE NEW ITEM"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func GetItemHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "GET {id} ITEMS"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func UpdateItemHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "UPDATE ITEM"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func DeleteItemHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "DELETE ITEM"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func UploadItemImageHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "UPLOAD ITEM'S IMAGE"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func LoginHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "LOGIN"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}

func LogoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getItems := "LOGOUT"
		fmt.Fprintf(w, "ANSWER: %15v", getItems)
	})
}
