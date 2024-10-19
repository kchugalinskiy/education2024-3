package main

import (
	"fmt"
	"github.com/kchugalinskiy/education2024-3/internal/echoservice"
	"net/http"

	"github.com/kchugalinskiy/education2024-3/internal/logic"
)

func main() {
	logic.ParGo()
	fmt.Println("Hello, World!")

	svc := echoservice.NewService()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: svc,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
