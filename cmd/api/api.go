package api

import (
	"net/http"

	"github.com/xan-mortum/api/internal/hello/router"
)

func main() {
	mux := http.NewServeMux()

	hw := router.NewRouter()
	mux.Handle("/api/v1/hello", hw.Get("/api/v1/hello"))

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}
