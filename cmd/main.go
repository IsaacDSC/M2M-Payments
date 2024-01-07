package main

import (
	"fmt"
	"log"
	"m2m/external"
	"m2m/external/wrap"
	"m2m/internal/container"
	"m2m/internal/shared"
	"net/http"
	"os"
	"strings"
)

func init() {
	if err := os.Mkdir("./tmp", os.ModePerm); err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			log.Fatal(err)
		}
	}
}

func main() {
	logger := wrap.NewInternalLogger(external.GeZapLogger())
	startServerHttp(logger)
}

func routers(logger shared.Logger) *http.ServeMux {
	r := http.NewServeMux()
	handlers := container.NewHandlers(logger).GetHandlers()
	for path, handler := range handlers {
		r.HandleFunc(string(path), handler)
	}
	return r
}

func startServerHttp(logger shared.Logger) {
	routers := routers(logger)
	logger.Info("Starting application")
	fmt.Println("Running in port 3000")
	http.ListenAndServe(":3000", routers)
}
