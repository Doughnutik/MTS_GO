package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func StartHandle(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Test string `json:"text"`
	}

	b, _ := io.ReadAll(r.Body)

	var rb requestBody

	_ = json.Unmarshal(b, &rb)

	fmt.Println(rb.Test)
	w.Write([]byte("Start our process"))
}

func main() {
	// Запускаем сервер на localhost с портом 8888. Чтобы что-то получить, можно отправить запрос curl http://localhost:8888/start
	mux := http.NewServeMux()
	mux.HandleFunc("/start", StartHandle)

	httpServer := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
	}
}
