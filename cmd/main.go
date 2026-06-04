package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ilydyu/opa_test/internal/auth"
	"github.com/ilydyu/opa_test/internal/controller"
	"github.com/ilydyu/opa_test/internal/middlware"
	"github.com/joho/godotenv"
)

// Для простоты опущены детали
func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	err := auth.InitOPA(os.Getenv("POLICY_PATH"))
	port := os.Getenv("PORT")

	if err != nil {
		log.Fatal(err)
	}

	controller := controller.NewController()
	middlware := middlware.NewMiddlware()

	mux.Handle("GET /resource", middlware.AuthMiddleware(http.HandlerFunc(controller.GetResource)))
	mux.Handle("POST /resource", middlware.AuthMiddleware(http.HandlerFunc(controller.CreateResource)))

	fmt.Printf("Server start on %s\n", port)

	err = http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
