package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fasthttp/router"

	"github.com/valyala/fasthttp"
)

func RecoverMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if e := recover(); e != nil {
				ctx.SetStatusCode(http.StatusInternalServerError)
				fmt.Fprintln(ctx, "we've got panic here!")
			}
		}()
		next(ctx)
	})
}
func LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		next(ctx)
		log.Printf("request: %s %s - %v\n",
			ctx.Method(),
			ctx.Path(),
			time.Since(start),
		)
	})
}
func main() {
	// как обычно, объявляем экземпляр роутера
	router := router.New()
	// регистрируем анонимную функцию-обработчик на корневой маршрут для метода GET
	// \n на конце строки добавлен, т.к. ctx.WriteString, в отличии от fmt.Fprintln не приклеевает перенос
	//строки
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("GET HANDLER\n")
	})
	// из-за того, что URI-параметры здесь хранятся в общем k-v контексте
	// нам необходимо приводить пустой интерфейс, получаемый по ключу id
	// к конкретному типу string и обязательно проверять, что у нас это получилось (иначе паника)
	router.GET("/{id}", func(ctx *fasthttp.RequestCtx) {
		id, ok := ctx.UserValue("id").(string)
		if !ok {
			ctx.WriteString("INVALID RESOURCE ID\n")
			// выполняет роль http.ResponseWriter.WriteHeader при записи статус-кода ответа
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
		// т.к. fasthttp.RequestCtx реализует io.Writer, мы можем для записи ответа с шаблоном уже
		//привычный нам fmt.Fprintf
		fmt.Fprintf(ctx, "GET BY ID HANDLER. RESOURCE ID IS %s\n", id)
	})

	log.Fatal(fasthttp.ListenAndServe(":8020",
		LoggingMiddleware(
			RecoverMiddleware(
				router.Handler,
			),
		),
	))
}
